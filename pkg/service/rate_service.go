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

func (s *rateService) LoadAndCacheAllRates() (map[string]model.Rate, error) {
	totalBuy := make(map[string]float64)
	totalSell := make(map[string]float64)
	counts := make(map[string]int)

	for _, adapter := range s.adapters {
		adapterRates, err := adapter.GetRates()
		if err == nil {
			for key, rate := range adapterRates {
				totalBuy[key] += rate.Buy
				totalSell[key] += rate.Sell
				counts[key]++
			}
		}
	}

	averageRates := make(map[string]model.Rate)
	for key, count := range counts {
		if count > 0 {
			averageRates[key] = model.Rate{
				Buy:  totalBuy[key] / float64(count),
				Sell: totalSell[key] / float64(count),
			}
		}
	}

	err := s.cachingRepo.SetRates(averageRates)
	if err != nil {
		fmt.Println("Failed to cache rates:", err)
	}

	return averageRates, nil
}

func (s *rateService) GetRateByType(rateType string) (*model.Rate, error) {
	cachedRates, err := s.cachingRepo.GetRates()
	if err == nil && cachedRates != nil {
		if time.Since(cachedRates.CreatedAt) < s.cacheTTLMinutes {
			if rate, exists := cachedRates.Rates[rateType]; exists {
				return &rate, nil
			}
		}
	}

	rates, err := s.LoadAndCacheAllRates()
	if err != nil {
		return nil, err
	}

	if rate, exists := rates[rateType]; exists {
		return &rate, nil
	}

	return nil, fmt.Errorf("rate type %s not found", rateType)
}
