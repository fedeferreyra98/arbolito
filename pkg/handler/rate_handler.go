package handler

import (
	"arbolito/pkg/service"
	"encoding/json"
	"net/http"
)

type RateHandler struct {
	service *service.RateService
}

func NewRateHandler(service *service.RateService) *RateHandler {
	return &RateHandler{service: service}
}

func (h *RateHandler) GetAverageRate(w http.ResponseWriter, r *http.Request) {
	rate, err := h.service.GetAverageRate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}
