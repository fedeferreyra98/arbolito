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
		w.Write([]byte(`{"bid": 140, "ask": 150}`))
	}))
	defer server.Close()

	api := NewCriptoyaAdapter(server.URL)
	rate, err := api.GetRate()

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 140.0, rate.Buy)
	assert.Equal(t, 150.0, rate.Sell)
}

func TestCriptoyaAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewCriptoyaAdapter(server.URL)
	rate, err := api.GetRate()

	assert.Error(t, err)
	assert.Nil(t, rate)
}
