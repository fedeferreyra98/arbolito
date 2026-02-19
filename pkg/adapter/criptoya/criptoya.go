package criptoya

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

type criptoyaAdapter struct {
	URL string
}

func NewCriptoyaAdapter(url string) repository.RateApiAdapter {
	return &criptoyaAdapter{URL: url}
}

func (c *criptoyaAdapter) GetRates() (map[string]model.Rate, error) {
	log.Printf("Fetching rates from Criptoya: %s", c.URL)
	resp, err := http.Get(c.URL)
	if err != nil {
		log.Printf("Error fetching from Criptoya: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding Criptoya response: %v", err)
		return nil, err
	}

	rates := make(map[string]model.Rate)

	extractSimple := func(key string) {
		if rateData, ok := data[key].(map[string]interface{}); ok {
			ask, _ := rateData["ask"].(float64)
			bid, _ := rateData["bid"].(float64)
			if ask == 0 && bid == 0 {
				price, _ := rateData["price"].(float64)
				ask = price
				bid = price
			}
			if ask > 0 || bid > 0 {
				rates[key] = model.Rate{Buy: bid, Sell: ask}
			}
		}
	}

	extractSimple("blue")
	extractSimple("oficial")
	extractSimple("tarjeta")

	extractComplex := func(key string) {
		if rateData, ok := data[key].(map[string]interface{}); ok {
			if al30, ok := rateData["al30"].(map[string]interface{}); ok {
				if d24hs, ok := al30["24hs"].(map[string]interface{}); ok {
					if price, ok := d24hs["price"].(float64); ok {
						rates[key] = model.Rate{Buy: price, Sell: price}
					}
				}
			}
		}
	}

	extractComplex("mep")
	extractComplex("ccl")

	log.Printf("Successfully fetched rates from Criptoya")
	return rates, nil
}
