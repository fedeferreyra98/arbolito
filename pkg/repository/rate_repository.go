package repository

import "arbolito/pkg/model"

type RateRepository interface {
	GetRate() (*model.Rate, error)
}
