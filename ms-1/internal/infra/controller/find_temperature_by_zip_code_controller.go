package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/entity"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/adapter"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/usecase"
	"go.opentelemetry.io/otel/baggage"
)

type findTemperatureByZipCodeController struct {
	Adapter entity.AdapterInterface
}

func NewFindTemperatureByZipCodeController() *findTemperatureByZipCodeController {
	adapter := adapter.NewMs2()
	return &findTemperatureByZipCodeController{Adapter: adapter}
}

func (f *findTemperatureByZipCodeController) FindTemperature(w http.ResponseWriter, r *http.Request) {

	var dto usecase.ZipCodeInputDto
	dto.ZipCode = r.URL.Query().Get("zipcode")
	if dto.ZipCode == "" {
		http.Error(w, "invalid zip code", http.StatusBadRequest)
		return
	}

	ctx := baggage.ContextWithoutBaggage(r.Context())

	findTemperatureByZip := usecase.NewTemperatureByZipCode(ctx, f.Adapter)
	response, err := findTemperatureByZip.Execute(ctx, dto)

	if err != nil {
		if err.Error() == "invalid zip code" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if err.Error() == "can not find zip code" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
