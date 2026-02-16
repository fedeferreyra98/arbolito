package service

import (
	"arbolito/pkg/mocks"
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRateService_GetAverageRate_FromCache(t *testing.T) {
	mockCachingRepo := new(mocks.MockCachingRepository)
	cachedRate := &model.CachedRate{
		Rate:      model.Rate{Buy: 100, Sell: 110},
		CreatedAt: time.Now(),
	}
	mockCachingRepo.On("GetRate").Return(cachedRate, nil)

	service := NewRateService(nil, mockCachingRepo)
	rate, err := service.GetAverageRate()

	assert.NoError(t, err)
	assert.Equal(t, cachedRate.Rate, *rate)
	mockCachingRepo.AssertExpectations(t)
}

func TestRateService_GetAverageRate_FromAPI(t *testing.T) {
	mockRateRepo1 := new(mocks.MockRateRepository)
	mockRateRepo2 := new(mocks.MockRateRepository)
	mockCachingRepo := new(mocks.MockCachingRepository)

	repos := []repository.RateApiAdapter{mockRateRepo1, mockRateRepo2}

	mockRateRepo1.On("GetRate").Return(&model.Rate{Buy: 100, Sell: 110}, nil)
	mockRateRepo2.On("GetRate").Return(&model.Rate{Buy: 102, Sell: 112}, nil)
	mockCachingRepo.On("GetRate").Return(nil, nil)
	mockCachingRepo.On("SetRate", mock.Anything).Return(nil)

	service := NewRateService(repos, mockCachingRepo)
	rate, err := service.GetAverageRate()

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 101.0, rate.Buy)
	assert.Equal(t, 111.0, rate.Sell)

	mockRateRepo1.AssertExpectations(t)
	mockRateRepo2.AssertExpectations(t)
	mockCachingRepo.AssertExpectations(t)
}

func TestRateService_GetAverageRate_CacheExpired(t *testing.T) {
	mockRateRepo := new(mocks.MockRateRepository)
	mockCachingRepo := new(mocks.MockCachingRepository)

	repos := []repository.RateApiAdapter{mockRateRepo}

	cachedRate := &model.CachedRate{
		Rate:      model.Rate{Buy: 100, Sell: 110},
		CreatedAt: time.Now().Add(-20 * time.Minute),
	}

	mockCachingRepo.On("GetRate").Return(cachedRate, nil)
	mockRateRepo.On("GetRate").Return(&model.Rate{Buy: 120, Sell: 130}, nil)
	mockCachingRepo.On("SetRate", mock.Anything).Return(nil)

	service := NewRateService(repos, mockCachingRepo)
	rate, err := service.GetAverageRate()

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 120.0, rate.Buy)
	assert.Equal(t, 130.0, rate.Sell)

	mockRateRepo.AssertExpectations(t)
	mockCachingRepo.AssertExpectations(t)
}
