package bluelytics

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"encoding/json"
	"net/http"
)

type bluelyticsRepository struct{
	URL string
}

func NewBluelyticsRepository(url string) repository.RateRepository {
	return &bluelyticsRepository{URL: url}
}

func (b *bluelyticsRepository) GetRate() (*model.Rate, error) {
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
