package mocks

import (
	"arbolito/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockCachingRepository struct {
	mock.Mock
}

func (m *MockCachingRepository) GetRate() (*model.CachedRate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.CachedRate), args.Error(1)
}

func (m *MockCachingRepository) SetRate(rate *model.Rate) error {
	args := m.Called(rate)
	return args.Error(0)
}
