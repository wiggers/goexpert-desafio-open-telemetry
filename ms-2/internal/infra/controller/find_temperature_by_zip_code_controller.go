package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/adapter"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/usecase"
	lib "github.com/wiggers/goexpert/desafio/1-temperatura/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type findTemperatureByZipCodeController struct {
	City    entity.CityAdapterInterface
	Weather entity.WeatherAdapterInterface
}

func NewFindTemperatureByZipCodeController() *findTemperatureByZipCodeController {

	city := adapter.NewViaCepApiData()
	weather := adapter.NewWeatherApi()

	return &findTemperatureByZipCodeController{
		City:    city,
		Weather: weather,
	}
}

func (f *findTemperatureByZipCodeController) FindTemperature(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	var dto usecase.ZipCodeInputDto
	dto.ZipCode = r.URL.Query().Get("zipcode")
	if dto.ZipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(lib.Response{Code: http.StatusBadRequest, Message: "Invalid zip code"})
		return
	}

	findTemperatureByZip := usecase.NewTemperatureByZipCode(ctx, f.City, f.Weather)
	response, err := findTemperatureByZip.Execute(dto)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(lib.Response{Code: err.Code, Message: err.Message})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lib.Response{Code: http.StatusOK, Message: response})
}
