package repository

import "arbolito/pkg/model"

type CachingRepository interface {
	GetRate() (*model.CachedRate, error)
	SetRate(*model.Rate) error
}
