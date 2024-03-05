package entity

import (
	"context"
)

type CityAdapterInterface interface {
	FindCity(ctx context.Context, zipcode *ZipCode) (*City, error)
}

type WeatherAdapterInterface interface {
	FindWeather(ctx context.Context, city *City) (*Weather, error)
}
