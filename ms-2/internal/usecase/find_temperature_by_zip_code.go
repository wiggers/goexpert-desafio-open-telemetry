package usecase

import (
	"context"
	"errors"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	"go.opentelemetry.io/otel"
)

type ZipCodeInputDto struct {
	ZipCode string `json:"cep"`
}

type ZipCodeOutputDto struct {
	City   string  `json:"city"`
	Temp_C float32 `json:"temp_C"`
	Temp_F float32 `json:"temp_F"`
	Temp_K float32 `json:"temp_K"`
}

type TemperatureByZipCode struct {
	Ctx            context.Context
	CityAdapter    entity.CityAdapterInterface
	WeatherAdapter entity.WeatherAdapterInterface
}

func NewTemperatureByZipCode(ctx context.Context, CityAdapter entity.CityAdapterInterface, WeatherAdapter entity.WeatherAdapterInterface) *TemperatureByZipCode {
	return &TemperatureByZipCode{
		Ctx:            ctx,
		CityAdapter:    CityAdapter,
		WeatherAdapter: WeatherAdapter,
	}
}

func (temp *TemperatureByZipCode) Execute(input ZipCodeInputDto) (ZipCodeOutputDto, error) {

	zipcode, err := entity.NewZipCode(input.ZipCode)
	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	tracer := otel.Tracer("ms2")
	ctx, span := tracer.Start(temp.Ctx, "find-city")
	city, err := temp.CityAdapter.FindCity(ctx, &zipcode)
	span.End()
	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	if !city.Exist() {
		return ZipCodeOutputDto{}, errors.New("can not find zip code")
	}

	ctx, span1 := tracer.Start(ctx, "find-weather")
	weather, err := temp.WeatherAdapter.FindWeather(ctx, city)
	span1.End()

	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	resul := ZipCodeOutputDto{
		City:   city.City,
		Temp_C: weather.Temp_C,
		Temp_F: weather.GetFahrenheit(),
		Temp_K: weather.GetKelvin(),
	}

	return resul, nil
}
