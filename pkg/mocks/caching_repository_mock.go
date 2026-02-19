package mocks

import (
	"arbolito/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockCachingRepository struct {
	mock.Mock
}

func (m *MockCachingRepository) GetRates() (*model.CachedRates, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.CachedRates), args.Error(1)
}

func (m *MockCachingRepository) SetRates(rates map[string]model.Rate) error {
	args := m.Called(rates)
	return args.Error(0)
}
