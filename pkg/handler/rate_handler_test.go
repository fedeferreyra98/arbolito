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
