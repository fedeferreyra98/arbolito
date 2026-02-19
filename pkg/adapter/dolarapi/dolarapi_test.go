package dolarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDolarAPI_GetRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"casa": "blue", "compra": 100, "venta": 110}, {"casa": "oficial", "compra": 90, "venta": 95}]`))
	}))
	defer server.Close()

	api := NewDolarAPIAdapter(server.URL)
	rates, err := api.GetRates()

	assert.NoError(t, err)
	assert.NotNil(t, rates)
	assert.Equal(t, 100.0, rates["blue"].Buy)
	assert.Equal(t, 110.0, rates["blue"].Sell)
	assert.Equal(t, 90.0, rates["oficial"].Buy)
	assert.Equal(t, 95.0, rates["oficial"].Sell)
}

func TestDolarAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewDolarAPIAdapter(server.URL)
	rates, err := api.GetRates()

	assert.Error(t, err)
	assert.Nil(t, rates)
}
