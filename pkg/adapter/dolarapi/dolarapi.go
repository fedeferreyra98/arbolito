package dolarapi

import (
	"arbolito/pkg/model"
	"encoding/json"
	"net/http"
)

type DolarAPI struct{
	URL string
}

func NewDolarAPI(url string) *DolarAPI {
	return &DolarAPI{URL: url}
}

func (d *DolarAPI) GetRate() (*model.Rate, error) {
	resp, err := http.Get(d.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Compra float64 `json:"compra"`
		Venta  float64 `json:"venta"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Rate{
		Buy:  data.Compra,
		Sell: data.Venta,
	}, nil
}
