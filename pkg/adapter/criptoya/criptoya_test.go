package criptoya

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriptoyaAPI_GetRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"blue": {"bid": 140, "ask": 150}, "oficial": {"bid": 100, "ask": 110}, "tarjeta": {"price": 180}, "mep": {"al30": {"24hs": {"price": 130}}}, "ccl": {"al30": {"24hs": {"price": 120}}}}`))
	}))
	defer server.Close()

	api := NewCriptoyaAdapter(server.URL)
	rates, err := api.GetRates()

	assert.NoError(t, err)
	assert.NotNil(t, rates)
	assert.Equal(t, 140.0, rates["blue"].Buy)
	assert.Equal(t, 150.0, rates["blue"].Sell)
	assert.Equal(t, 100.0, rates["oficial"].Buy)
	assert.Equal(t, 110.0, rates["oficial"].Sell)
	assert.Equal(t, 180.0, rates["tarjeta"].Buy)
	assert.Equal(t, 130.0, rates["mep"].Buy)
	assert.Equal(t, 120.0, rates["ccl"].Buy)
}

func TestCriptoyaAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewCriptoyaAdapter(server.URL)
	rates, err := api.GetRates()

	assert.Error(t, err)
	assert.Nil(t, rates)
}
