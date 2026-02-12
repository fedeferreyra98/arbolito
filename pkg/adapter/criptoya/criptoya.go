package criptoya

import (
	"arbolito/pkg/model"
	"encoding/json"
	"net/http"
)

type CriptoyaAPI struct {
	URL string
}

func NewCriptoyaAPI(url string) *CriptoyaAPI {
	return &CriptoyaAPI{URL: url}
}

func (c *CriptoyaAPI) GetRate() (*model.Rate, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Blue struct {
			Ask float64 `json:"ask"`
			Bid float64 `json:"bid"`
		} `json:"blue"`
	}

	// The response is a map of strings to interfaces, so we need to decode it into a map

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Rate{
		Buy:  data.Blue.Bid,
		Sell: data.Blue.Ask,
	}, nil
}
