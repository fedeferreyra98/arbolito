package service

import "arbolito/pkg/model"

type RateService interface {
	LoadAndCacheAllRates() (map[string]model.Rate, error)
	GetRateByType(rateType string) (*model.Rate, error)
}
