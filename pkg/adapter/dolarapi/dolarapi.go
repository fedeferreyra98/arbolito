package dolarapi

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"net/http"
)

type dolarapiAdapter struct {
	URL string
}

func NewDolarAPIAdapter(url string) repository.RateApiAdapter {
	return &dolarapiAdapter{URL: url}
}

func (d *dolarapiAdapter) GetRate() (*model.Rate, error) {
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
