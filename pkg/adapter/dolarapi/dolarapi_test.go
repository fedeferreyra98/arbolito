package dolarapi

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDolarAPI_GetRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"compra": 100, "venta": 110}`))
	}))
	defer server.Close()

	api := NewDolarAPIRepository(server.URL)
	rate, err := api.GetRate()

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 100.0, rate.Buy)
	assert.Equal(t, 110.0, rate.Sell)
}

func TestDolarAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewDolarAPIRepository(server.URL)
	rate, err := api.GetRate()

	assert.Error(t, err)
	assert.Nil(t, rate)
}
