package mocks

import (
	"arbolito/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockRateService struct {
	mock.Mock
}

func (m *MockRateService) LoadAndCacheAllRates() (map[string]model.Rate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]model.Rate), args.Error(1)
}

func (m *MockRateService) GetRateByType(rateType string) (*model.Rate, error) {
	args := m.Called(rateType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Rate), args.Error(1)
}
