package bluelytics

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBluelyticsAPI_GetRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"blue": {"value_buy": 120, "value_sell": 130}}`))
	}))
	defer server.Close()

	api := NewBluelyticsRepository(server.URL)
	rate, err := api.GetRate()

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 120.0, rate.Buy)
	assert.Equal(t, 130.0, rate.Sell)
}

func TestBluelyticsAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := NewBluelyticsRepository(server.URL)
	rate, err := api.GetRate()

	assert.Error(t, err)
	assert.Nil(t, rate)
}
