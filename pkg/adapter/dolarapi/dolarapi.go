package dolarapi

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

type dolarapiAdapter struct {
	URL string
}

func NewDolarAPIAdapter(url string) repository.RateApiAdapter {
	return &dolarapiAdapter{URL: url}
}

func (d *dolarapiAdapter) GetRate() (*model.Rate, error) {
	log.Printf("Fetching rate from DolarAPI: %s", d.URL)
	resp, err := http.Get(d.URL)
	if err != nil {
		log.Printf("Error fetching from DolarAPI: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Compra float64 `json:"compra"`
		Venta  float64 `json:"venta"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding DolarAPI response: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched rate from DolarAPI")
	return &model.Rate{
		Buy:  data.Compra,
		Sell: data.Venta,
	}, nil
}
