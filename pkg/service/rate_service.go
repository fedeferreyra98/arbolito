package service

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
)

type RateService struct {
	repos []repository.RateRepository
}

func NewRateService(repos []repository.RateRepository) *RateService {
	return &RateService{repos: repos}
}

func (s *RateService) GetAverageRate() (*model.Rate, error) {
	var totalBuy, totalSell float64
	var count int

	for _, repo := range s.repos {
		rate, err := repo.GetRate()
		if err == nil {
			totalBuy += rate.Buy
			totalSell += rate.Sell
			count++
		}
	}

	if count == 0 {
		return &model.Rate{
			Buy:  0,
			Sell: 0,
		}, nil
	}

	return &model.Rate{
		Buy:  totalBuy / float64(count),
		Sell: totalSell / float64(count),
	}, nil
}
