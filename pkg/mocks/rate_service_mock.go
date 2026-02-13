package mocks

import (
	"arbolito/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockRateService struct {
	mock.Mock
}

func (m *MockRateService) GetAverageRate() (*model.Rate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Rate), args.Error(1)
}
