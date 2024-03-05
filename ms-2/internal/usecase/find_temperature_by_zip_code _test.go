package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
)

type CityAdapterMock struct {
	mock.Mock
}

type WeatherAdapterMock struct {
	mock.Mock
}

func (adapter *CityAdapterMock) FindCity(ctx context.Context, zipcode *entity.ZipCode) (*entity.City, error) {
	args := adapter.Called(ctx, zipcode)
	return args.Get(0).(*entity.City), args.Error(1)
}

func (adapter *WeatherAdapterMock) FindWeather(ctx context.Context, city *entity.City) (*entity.Weather, error) {
	args := adapter.Called(ctx, city)
	return args.Get(0).(*entity.Weather), args.Error(1)
}

func TestGivenExistingZipCode_ThenShouldReceivedTemperatureWithParameters(t *testing.T) {
	cityAdapter := CityAdapterMock{}
	weatherAdapter := WeatherAdapterMock{}
	var zipcode ZipCodeInputDto
	zipcode.ZipCode = "88807278"
	cityAdapter.On("FindCity", mock.MatchedBy(func(context.Context) bool { return true }), &entity.ZipCode{Cep: "88807278"}).Return(&entity.City{City: "Joinville"}, nil)
	weatherAdapter.On("FindWeather", mock.MatchedBy(func(context.Context) bool { return true }), &entity.City{City: "Joinville"}).Return(&entity.Weather{Temp_C: 29}, nil)

	findTemperatureByZip := NewTemperatureByZipCode(&cityAdapter, &weatherAdapter)
	response, err := findTemperatureByZip.Execute(zipcode)
	assert.Nil(t, err)
	assert.Equal(t, response.Temp_C, float32(29))
	assert.Equal(t, response.Temp_F, float32(84.2))
	assert.Equal(t, response.Temp_K, float32(302))
}

func TestGivenAInvalidZipCode_ThenShouldReceiveAnError(t *testing.T) {
	cityAdapter := CityAdapterMock{}
	weatherAdapter := WeatherAdapterMock{}
	var zipcode ZipCodeInputDto
	zipcode.ZipCode = "AAAAAAAAA"

	findTemperatureByZip := NewTemperatureByZipCode(&cityAdapter, &weatherAdapter)
	response, err := findTemperatureByZip.Execute(zipcode)
	assert.Error(t, err, " invalid zipcode")
	assert.Equal(t, response.Temp_C, float32(0.0))
	assert.Equal(t, response.Temp_F, float32(0.0))
	assert.Equal(t, response.Temp_K, float32(0.0))
}

func TestGivenAZipCodeNotExisting_ThenShouldReceiveAnError(t *testing.T) {
	cityAdapter := CityAdapterMock{}
	weatherAdapter := WeatherAdapterMock{}
	var zipcode ZipCodeInputDto
	zipcode.ZipCode = "99999999"
	cityAdapter.On("FindCity", mock.MatchedBy(func(context.Context) bool { return true }), &entity.ZipCode{Cep: "99999999"}).Return(&entity.City{City: ""}, nil)

	findTemperatureByZip := NewTemperatureByZipCode(&cityAdapter, &weatherAdapter)
	response, err := findTemperatureByZip.Execute(zipcode)
	assert.Error(t, err, "can not find zip code")
	assert.Equal(t, response.Temp_C, float32(0.0))
	assert.Equal(t, response.Temp_F, float32(0.0))
	assert.Equal(t, response.Temp_K, float32(0.0))
}
