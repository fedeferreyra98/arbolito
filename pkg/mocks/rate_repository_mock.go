package mocks

import (
	"arbolito/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockRateRepository struct {
	mock.Mock
}

func (m *MockRateRepository) GetRates() (map[string]model.Rate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]model.Rate), args.Error(1)
}
