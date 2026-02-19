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

// GetBlueRate godoc
// @Summary      Get blue dollar rate
// @Description  Get the blue dollar rate aggregated from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-blue [get]
func (h *RateHandler) GetBlueRate(w http.ResponseWriter, r *http.Request) {
	h.handleRateRequest(w, "blue")
}

// GetOficialRate godoc
// @Summary      Get oficial dollar rate
// @Description  Get the oficial dollar rate aggregated from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-oficial [get]
func (h *RateHandler) GetOficialRate(w http.ResponseWriter, r *http.Request) {
	h.handleRateRequest(w, "oficial")
}

// GetMepRate godoc
// @Summary      Get MEP dollar rate
// @Description  Get the MEP dollar rate aggregated from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-mep [get]
func (h *RateHandler) GetMepRate(w http.ResponseWriter, r *http.Request) {
	h.handleRateRequest(w, "mep")
}

// GetCclRate godoc
// @Summary      Get CCL dollar rate
// @Description  Get the CCL dollar rate aggregated from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-ccl [get]
func (h *RateHandler) GetCclRate(w http.ResponseWriter, r *http.Request) {
	h.handleRateRequest(w, "ccl")
}

// GetTarjetaRate godoc
// @Summary      Get Tarjeta dollar rate
// @Description  Get the tarjeta dollar rate aggregated from different sources
// @Tags         rates
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Rate
// @Router       /dolar-tarjeta [get]
func (h *RateHandler) GetTarjetaRate(w http.ResponseWriter, r *http.Request) {
	h.handleRateRequest(w, "tarjeta")
}

func (h *RateHandler) handleRateRequest(w http.ResponseWriter, rateType string) {
	rate, err := h.service.GetRateByType(rateType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}
