package bluelytics

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

type bluelyticsAdapter struct {
	URL string
}

func NewBluelyticsAdapter(url string) repository.RateApiAdapter {
	return &bluelyticsAdapter{URL: url}
}

func (b *bluelyticsAdapter) GetRate() (*model.Rate, error) {
	log.Printf("Fetching rate from Bluelytics: %s", b.URL)
	resp, err := http.Get(b.URL)
	if err != nil {
		log.Printf("Error fetching from Bluelytics: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Blue struct {
			ValueAvg  float64 `json:"value_avg"`
			ValueSell float64 `json:"value_sell"`
			ValueBuy  float64 `json:"value_buy"`
		} `json:"blue"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding Bluelytics response: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched rate from Bluelytics")
	return &model.Rate{
		Buy:  data.Blue.ValueBuy,
		Sell: data.Blue.ValueSell,
	}, nil
}
