package criptoya

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type criptoyaAdapter struct {
	URL string
}

func NewCriptoyaAdapter(url string) repository.RateRepository {
	return &criptoyaAdapter{URL: url}
}

func (c *criptoyaAdapter) GetRate() (*model.Rate, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Assert the type of the values to float64
	ask, ok := data["ask"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to assert ask to float64")
	}

	bid, ok := data["bid"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to assert bid to float64")
	}

	return &model.Rate{
		Buy:  bid,
		Sell: ask,
	}, nil
}
