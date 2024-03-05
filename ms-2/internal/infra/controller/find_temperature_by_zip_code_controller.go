package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/adapter"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/usecase"
	"go.opentelemetry.io/otel/baggage"
)

type findTemperatureByZipCodeController struct {
	City    entity.CityAdapterInterface
	Weather entity.WeatherAdapterInterface
}

type sendError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewFindTemperatureByZipCodeController() *findTemperatureByZipCodeController {

	city := adapter.NewBrasilApiData()
	weather := adapter.NewWeatherApi()

	return &findTemperatureByZipCodeController{
		City:    city,
		Weather: weather,
	}
}

func (f *findTemperatureByZipCodeController) FindTemperature(w http.ResponseWriter, r *http.Request) {

	ctx := baggage.ContextWithoutBaggage(r.Context())

	var dto usecase.ZipCodeInputDto
	dto.ZipCode = r.URL.Query().Get("zipcode")
	if dto.ZipCode == "" {
		http.Error(w, "invalid zip code", http.StatusBadRequest)
		return
	}

	findTemperatureByZip := usecase.NewTemperatureByZipCode(ctx, f.City, f.Weather)

	response, err := findTemperatureByZip.Execute(dto)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		if err.Error() == "invalid zip code" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(sendError{Message: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		if err.Error() == "can not find zip code" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(sendError{Message: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		json.NewEncoder(w).Encode(sendError{Message: err.Error(), Code: http.StatusInternalServerError})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
