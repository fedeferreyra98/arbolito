package repository

import "arbolito/pkg/model"

type RateApiAdapter interface {
	GetRate() (*model.Rate, error)
}
