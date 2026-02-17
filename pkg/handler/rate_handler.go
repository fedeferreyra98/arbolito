package handler

import (
	"arbolito/pkg/service"
	"encoding/json"
	"net/http"
)

type RateHandler struct {
	service service.RateService
}

func NewRateHandler(service service.RateService) *RateHandler {
	return &RateHandler{service: service}
}

// GetAverageRate godoc
// @Summary      Get average dollar rate
// @Description  Get the average dollar rate from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-blue [get]
func (h *RateHandler) GetAverageRate(w http.ResponseWriter, r *http.Request) {
	rate, err := h.service.GetAverageRate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}
