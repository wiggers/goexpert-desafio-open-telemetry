package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/entity"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/adapter"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/usecase"
	lib "github.com/wiggers/goexpert/desafio/temperature-ms-1/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
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
	err := json.NewDecoder(r.Body).Decode(&dto)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(lib.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if dto.ZipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(lib.Response{Code: http.StatusBadRequest, Message: "Check you zipcode"})
		return
	}

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	findTemperatureByZip := usecase.NewTemperatureByZipCode(ctx, f.Adapter)
	response, errorUseCase := findTemperatureByZip.Execute(ctx, dto)

	if errorUseCase != nil {
		w.WriteHeader(errorUseCase.Code)
		json.NewEncoder(w).Encode(lib.Response{Code: errorUseCase.Code, Message: errorUseCase.Message})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lib.Response{Code: http.StatusOK, Message: response})
}
