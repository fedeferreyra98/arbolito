package handler

import (
	"arbolito/pkg/model"
	"arbolito/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type RateHandler struct {
	service service.RateService
}

func NewRateHandler(service service.RateService) *RateHandler {
	return &RateHandler{service: service}
}

// Convert godoc
// @Summary      Convert amount between ARS and USD
// @Description  Convert a specified amount from ARS to USD or USD to ARS using a given rate type
// @Tags         rates
// @Accept       json
// @Produce      json
// @Param        amount    query     number  true  "Amount to convert"
// @Param        from      query     string  true  "Currency to convert from (ARS, USD)"
// @Param        to        query     string  true  "Currency to convert to (ARS, USD)"
// @Param        rate_type query     string  true  "Type of dollar rate to use (e.g., blue, oficial)"
// @Success      200  {object}  model.ConversionResponse
// @Router       /convert [get]
func (h *RateHandler) Convert(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	rateType := r.URL.Query().Get("rate_type")

	if amountStr == "" || from == "" || to == "" || rateType == "" {
		http.Error(w, "amount, from, to, and rate_type are required query parameters", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	if (from != "ARS" && from != "USD") || (to != "ARS" && to != "USD") || from == to {
		http.Error(w, "invalid currency conversion pair (must be ARS to USD or USD to ARS)", http.StatusBadRequest)
		return
	}

	rate, err := h.service.GetRateByType(rateType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var convertedAmount float64
	var rateApplied float64

	if from == "ARS" && to == "USD" {
		convertedAmount = amount / rate.Sell
		rateApplied = rate.Sell
	} else if from == "USD" && to == "ARS" {
		convertedAmount = amount * rate.Buy
		rateApplied = rate.Buy
	}

	response := model.ConversionResponse{
		Amount:          amount,
		ConvertedAmount: convertedAmount,
		From:            from,
		To:              to,
		RateType:        rateType,
		RateApplied:     rateApplied,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
