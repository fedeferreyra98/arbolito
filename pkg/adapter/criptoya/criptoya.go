package criptoya

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type criptoyaRepository struct {
	URL string
}

func NewCriptoyaRepository(url string) repository.RateRepository {
	return &criptoyaRepository{URL: url}
}

func (c *criptoyaRepository) GetRate() (*model.Rate, error) {
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
