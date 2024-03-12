package usecase

import (
	"context"
	"net/http"

	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/entity"
	lib "github.com/wiggers/goexpert/desafio/temperature-ms-1/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
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
	Adapter entity.AdapterInterface
	Ctx     context.Context
}

func NewTemperatureByZipCode(ctx context.Context, Adapter entity.AdapterInterface) *TemperatureByZipCode {
	return &TemperatureByZipCode{Ctx: ctx, Adapter: Adapter}
}

func (temp *TemperatureByZipCode) Execute(ctx context.Context, input ZipCodeInputDto) (ZipCodeOutputDto, *lib.AppError) {

	tracer := otel.Tracer("ms1")
	ctx, span := tracer.Start(ctx, "total-ms-1",
		trace.WithSpanKind(trace.SpanKindServer))

	zipcode, err := entity.NewZipCode(input.ZipCode)
	if err != nil {
		return ZipCodeOutputDto{}, &lib.AppError{ErrorMsg: err, Message: err.Error(), Code: http.StatusInternalServerError}
	}

	temperature, err := temp.Adapter.FindData(ctx, &zipcode)
	span.End()

	if err != nil {
		if err.Error() == "Invalid zip code" {
			return ZipCodeOutputDto{}, &lib.AppError{ErrorMsg: err, Message: err.Error(), Code: http.StatusUnprocessableEntity}
		}
		if err.Error() == "Can not find zipcode" {
			return ZipCodeOutputDto{}, &lib.AppError{ErrorMsg: err, Message: err.Error(), Code: http.StatusNotFound}
		}

		return ZipCodeOutputDto{}, &lib.AppError{ErrorMsg: err, Message: err.Error(), Code: http.StatusInternalServerError}
	}

	return ZipCodeOutputDto{City: temperature.City, Temp_C: temperature.Temp_C, Temp_F: temperature.Temp_F, Temp_K: temperature.Temp_K}, nil
}
