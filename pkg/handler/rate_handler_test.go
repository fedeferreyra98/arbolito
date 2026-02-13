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

func TestRateHandler_GetAverageRate(t *testing.T) {
	mockService := new(mocks.MockRateService)
	rate := &model.Rate{Buy: 100, Sell: 110}
	mockService.On("GetAverageRate").Return(rate, nil)

	handler := NewRateHandler(mockService)

	req, err := http.NewRequest("GET", "/dolar-blue", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAverageRate(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseRate model.Rate
	err = json.Unmarshal(rr.Body.Bytes(), &responseRate)
	assert.NoError(t, err)

	assert.Equal(t, *rate, responseRate)
	mockService.AssertExpectations(t)
}
