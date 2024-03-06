package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	lib "github.com/wiggers/goexpert/desafio/1-temperatura/pkg"
)

type currentWeather struct {
	Current currentData `json:"current"`
}

type currentData struct {
	Temp_C float32 `json:"temp_c"`
}

func NewWeatherApi() *currentWeather {
	return &currentWeather{}
}

func (c *currentWeather) FindWeather(ctx context.Context, city *entity.City) (*entity.Weather, error) {
	dataWeather, err := lib.CallHttpGetWeather(ctx, city.City)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	var current currentWeather
	json.Unmarshal(dataWeather, &current)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}
	return &entity.Weather{Temp_C: current.Current.Temp_C}, nil
}
