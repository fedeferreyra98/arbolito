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

func (d *dolarapiAdapter) GetRates() (map[string]model.Rate, error) {
	log.Printf("Fetching rates from DolarAPI: %s", d.URL)
	resp, err := http.Get(d.URL)
	if err != nil {
		log.Printf("Error fetching from DolarAPI: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data []struct {
		Casa   string  `json:"casa"`
		Compra float64 `json:"compra"`
		Venta  float64 `json:"venta"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding DolarAPI response: %v", err)
		return nil, err
	}

	rates := make(map[string]model.Rate)
	for _, item := range data {
		key := ""
		switch item.Casa {
		case "oficial":
			key = "oficial"
		case "blue":
			key = "blue"
		case "bolsa":
			key = "mep"
		case "contadoconliqui":
			key = "ccl"
		case "tarjeta":
			key = "tarjeta"
		}
		if key != "" {
			rates[key] = model.Rate{
				Buy:  item.Compra,
				Sell: item.Venta,
			}
		}
	}

	log.Printf("Successfully fetched rates from DolarAPI")
	return rates, nil
}
