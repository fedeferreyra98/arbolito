package repository

import "arbolito/pkg/model"

type CachingRepository interface {
	GetRates() (*model.CachedRates, error)
	SetRates(rates map[string]model.Rate) error
}
