package handler

import (
	"arbolito/pkg/mocks"
	"arbolito/pkg/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateHandler_GetRates(t *testing.T) {
	mockService := new(mocks.MockRateService)
	rate := &model.Rate{Buy: 100, Sell: 110}

	handler := NewRateHandler(mockService)

	tests := []struct {
		name       string
		methodFn   func(w http.ResponseWriter, r *http.Request)
		url        string
		targetType string
	}{
		{"GetBlueRate", handler.GetBlueRate, "/dolar-blue", "blue"},
		{"GetOficialRate", handler.GetOficialRate, "/dolar-oficial", "oficial"},
		{"GetMepRate", handler.GetMepRate, "/dolar-mep", "mep"},
		{"GetCclRate", handler.GetCclRate, "/dolar-ccl", "ccl"},
		{"GetTarjetaRate", handler.GetTarjetaRate, "/dolar-tarjeta", "tarjeta"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.ExpectedCalls = nil // Clear expected calls for the next run
			mockService.On("GetRateByType", tt.targetType).Return(rate, nil)

			req, err := http.NewRequest("GET", tt.url, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			tt.methodFn(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)

			var responseRate model.Rate
			err = json.Unmarshal(rr.Body.Bytes(), &responseRate)
			assert.NoError(t, err)

			assert.Equal(t, *rate, responseRate)
			mockService.AssertExpectations(t)
		})
	}
}

func TestRateHandler_Convert(t *testing.T) {
	mockService := new(mocks.MockRateService)
	// rate.Buy is 100, rate.Sell is 110
	rate := &model.Rate{Buy: 100, Sell: 110}

	handler := NewRateHandler(mockService)

	tests := []struct {
		name           string
		url            string
		mockRateType   string
		mockReturnRate *model.Rate
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ARS to USD",
			url:            "/convert?amount=1100&from=ARS&to=USD&rate_type=blue",
			mockRateType:   "blue",
			mockReturnRate: rate,
			expectedStatus: http.StatusOK,
			// 1100 / Sell (110) = 10
			expectedBody: `{"amount":1100,"converted_amount":10,"from":"ARS","to":"USD","rate_type":"blue","rate_applied":110}`,
		},
		{
			name:           "Valid USD to ARS",
			url:            "/convert?amount=10&from=USD&to=ARS&rate_type=blue",
			mockRateType:   "blue",
			mockReturnRate: rate,
			expectedStatus: http.StatusOK,
			// 10 * Buy (100) = 1000
			expectedBody: `{"amount":10,"converted_amount":1000,"from":"USD","to":"ARS","rate_type":"blue","rate_applied":100}`,
		},
		{
			name:           "Missing params",
			url:            "/convert?amount=10&from=USD",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "amount, from, to, and rate_type are required query parameters\n",
		},
		{
			name:           "Invalid amount",
			url:            "/convert?amount=-10&from=USD&to=ARS&rate_type=blue",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid amount\n",
		},
		{
			name:           "Invalid currency pair",
			url:            "/convert?amount=10&from=EUR&to=ARS&rate_type=blue",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid currency conversion pair (must be ARS to USD or USD to ARS)\n",
		},
		{
			name:           "Same currency",
			url:            "/convert?amount=10&from=USD&to=USD&rate_type=blue",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid currency conversion pair (must be ARS to USD or USD to ARS)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.ExpectedCalls = nil
			if tt.mockRateType != "" {
				mockService.On("GetRateByType", tt.mockRateType).Return(tt.mockReturnRate, tt.mockReturnErr)
			}

			req, err := http.NewRequest("GET", tt.url, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.Convert(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			mockService.AssertExpectations(t)
		})
	}
}
