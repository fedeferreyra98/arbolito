package criptoya

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type criptoyaAdapter struct {
	URL string
}

func NewCriptoyaAdapter(url string) repository.RateApiAdapter {
	return &criptoyaAdapter{URL: url}
}

func (c *criptoyaAdapter) GetRate() (*model.Rate, error) {
	log.Printf("Fetching rate from Criptoya: %s", c.URL)
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

	// Access the "blue" key
	blueData, ok := data["blue"].(map[string]interface{})
	if !ok {
		log.Printf("Failed to assert data['blue'] to map[string]interface{}")
		return nil, fmt.Errorf("failed to assert data['blue'] to map[string]interface{}")
	}

	// Assert the type of the values to float64
	ask, ok := blueData["ask"].(float64)
	if !ok {
		log.Printf("Failed to assert ask to float64")
		return nil, fmt.Errorf("failed to assert ask to float64")
	}

	bid, ok := blueData["bid"].(float64)
	if !ok {
		log.Printf("Failed to assert bid to float64")
		return nil, fmt.Errorf("failed to assert bid to float64")
	}

	log.Printf("Successfully fetched rate from Criptoya")
	return &model.Rate{
		Buy:  bid,
		Sell: ask,
	}, nil
}
