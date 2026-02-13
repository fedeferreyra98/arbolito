package service

import "arbolito/pkg/model"

type RateService interface {
	GetAverageRate() (*model.Rate, error)
}
