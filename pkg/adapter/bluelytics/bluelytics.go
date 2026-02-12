package bluelytics

import (
	"arbolito/pkg/model"
	"encoding/json"
	"net/http"
)

type BluelyticsAPI struct{
	URL string
}

func NewBluelyticsAPI(url string) *BluelyticsAPI {
	return &BluelyticsAPI{URL: url}
}

func (b *BluelyticsAPI) GetRate() (*model.Rate, error) {
	resp, err := http.Get(b.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Blue struct {
			ValueAvg   float64 `json:"value_avg"`
			ValueSell float64 `json:"value_sell"`
			ValueBuy  float64 `json:"value_buy"`
		} `json:"blue"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Rate{
		Buy:  data.Blue.ValueBuy,
		Sell: data.Blue.ValueSell,
	}, nil
}
