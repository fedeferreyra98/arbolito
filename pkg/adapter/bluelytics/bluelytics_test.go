package bluelytics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBluelyticsAPI_GetRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"blue": {"value_buy": 120, "value_sell": 130}, "oficial": {"value_buy": 100, "value_sell": 110}}`))
	}))
	defer server.Close()

	api := NewBluelyticsAdapter(server.URL)
	rates, err := api.GetRates()

	assert.NoError(t, err)
	assert.NotNil(t, rates)
	assert.Equal(t, 120.0, rates["blue"].Buy)
	assert.Equal(t, 130.0, rates["blue"].Sell)
	assert.Equal(t, 100.0, rates["oficial"].Buy)
	assert.Equal(t, 110.0, rates["oficial"].Sell)
}

func TestBluelyticsAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewBluelyticsAdapter(server.URL)
	rates, err := api.GetRates()

	assert.Error(t, err)
	assert.Nil(t, rates)
}
