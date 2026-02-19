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

func TestRateService_LoadAndCacheAllRates(t *testing.T) {
	mockRateRepo1 := new(mocks.MockRateRepository)
	mockRateRepo2 := new(mocks.MockRateRepository)
	mockCachingRepo := new(mocks.MockCachingRepository)

	repos := []repository.RateApiAdapter{mockRateRepo1, mockRateRepo2}

	mockRateRepo1.On("GetRates").Return(map[string]model.Rate{"blue": {Buy: 100, Sell: 110}}, nil)
	mockRateRepo2.On("GetRates").Return(map[string]model.Rate{"blue": {Buy: 102, Sell: 112}, "oficial": {Buy: 90, Sell: 95}}, nil)
	mockCachingRepo.On("SetRates", mock.Anything).Return(nil)

	service := NewRateService(repos, mockCachingRepo)
	rates, err := service.LoadAndCacheAllRates()

	assert.NoError(t, err)
	assert.NotNil(t, rates)
	assert.Equal(t, 101.0, rates["blue"].Buy)
	assert.Equal(t, 111.0, rates["blue"].Sell)
	assert.Equal(t, 90.0, rates["oficial"].Buy)
	assert.Equal(t, 95.0, rates["oficial"].Sell)

	mockRateRepo1.AssertExpectations(t)
	mockRateRepo2.AssertExpectations(t)
	mockCachingRepo.AssertExpectations(t)
}

func TestRateService_GetRateByType_FromCache(t *testing.T) {
	mockCachingRepo := new(mocks.MockCachingRepository)
	cachedRates := &model.CachedRates{
		Rates:     map[string]model.Rate{"blue": {Buy: 100, Sell: 110}},
		CreatedAt: time.Now(),
	}
	mockCachingRepo.On("GetRates").Return(cachedRates, nil)

	service := NewRateService(nil, mockCachingRepo)
	rate, err := service.GetRateByType("blue")

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 100.0, rate.Buy)
	assert.Equal(t, 110.0, rate.Sell)
	mockCachingRepo.AssertExpectations(t)
}

func TestRateService_GetRateByType_CacheExpired(t *testing.T) {
	mockRateRepo := new(mocks.MockRateRepository)
	mockCachingRepo := new(mocks.MockCachingRepository)

	repos := []repository.RateApiAdapter{mockRateRepo}

	cachedRates := &model.CachedRates{
		Rates:     map[string]model.Rate{"blue": {Buy: 100, Sell: 110}},
		CreatedAt: time.Now().Add(-20 * time.Minute),
	}

	mockCachingRepo.On("GetRates").Return(cachedRates, nil)
	mockRateRepo.On("GetRates").Return(map[string]model.Rate{"blue": {Buy: 120, Sell: 130}}, nil)
	mockCachingRepo.On("SetRates", mock.Anything).Return(nil)

	service := NewRateService(repos, mockCachingRepo)
	rate, err := service.GetRateByType("blue")

	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.Equal(t, 120.0, rate.Buy)
	assert.Equal(t, 130.0, rate.Sell)

	mockRateRepo.AssertExpectations(t)
	mockCachingRepo.AssertExpectations(t)
}
