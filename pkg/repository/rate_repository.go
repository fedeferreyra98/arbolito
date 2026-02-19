package repository

import "arbolito/pkg/model"

type RateApiAdapter interface {
	GetRates() (map[string]model.Rate, error)
}
