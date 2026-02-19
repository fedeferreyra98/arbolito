package model

type ConversionResponse struct {
	Amount          float64 `json:"amount"`
	ConvertedAmount float64 `json:"converted_amount"`
	From            string  `json:"from"`
	To              string  `json:"to"`
	RateType        string  `json:"rate_type"`
	RateApplied     float64 `json:"rate_applied"`
}
