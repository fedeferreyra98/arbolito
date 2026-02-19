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

func (b *bluelyticsAdapter) GetRates() (map[string]model.Rate, error) {
	log.Printf("Fetching rates from Bluelytics: %s", b.URL)
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
		Oficial struct {
			ValueAvg  float64 `json:"value_avg"`
			ValueSell float64 `json:"value_sell"`
			ValueBuy  float64 `json:"value_buy"`
		} `json:"oficial"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding Bluelytics response: %v", err)
		return nil, err
	}

	rates := make(map[string]model.Rate)
	rates["blue"] = model.Rate{
		Buy:  data.Blue.ValueBuy,
		Sell: data.Blue.ValueSell,
	}
	rates["oficial"] = model.Rate{
		Buy:  data.Oficial.ValueBuy,
		Sell: data.Oficial.ValueSell,
	}

	log.Printf("Successfully fetched rates from Bluelytics")
	return rates, nil
}
