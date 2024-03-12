package adapter

import (
	"context"
	"encoding/json"
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

	params := map[string]string{"key": "2345b10c20c34affa4c170312241602", "q": city.City}
	dataWeather, err := lib.CallHttpGet(ctx, "https://api.weatherapi.com/v1/current.json", params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var current currentWeather
	err = json.Unmarshal(dataWeather, &current)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &entity.Weather{Temp_C: current.Current.Temp_C}, nil
}
