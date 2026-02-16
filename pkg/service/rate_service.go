package service

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"fmt"
	"time"
)

type rateService struct {
	adapters        []repository.RateApiAdapter
	cachingRepo     repository.CachingRepository
	cacheTTLMinutes time.Duration
}

func NewRateService(adapters []repository.RateApiAdapter, cachingRepo repository.CachingRepository) RateService {
	return &rateService{
		adapters:        adapters,
		cachingRepo:     cachingRepo,
		cacheTTLMinutes: 15 * time.Minute,
	}
}

func (s *rateService) GetAverageRate() (*model.Rate, error) {
	// Try to get from cache first
	cachedRate, err := s.cachingRepo.GetRate()
	if err == nil && cachedRate != nil {
		if time.Since(cachedRate.CreatedAt) < s.cacheTTLMinutes {
			return &cachedRate.Rate, nil
		}
	}

	// If not in cache or expired, fetch from APIs
	var totalBuy, totalSell float64
	var count int

	for _, adapter := range s.adapters {
		rate, err := adapter.GetRate()
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

	averageRate := &model.Rate{
		Buy:  totalBuy / float64(count),
		Sell: totalSell / float64(count),
	}

	// Save to cache
	err = s.cachingRepo.SetRate(averageRate)
	if err != nil {
		fmt.Println("Failed to cache rate:", err)
	}

	return averageRate, nil
}
