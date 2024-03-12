package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Ms2 struct {
	Success Success
	Error   Error
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Message Body `json:"message"`
}

type Body struct {
	City   string  `json:"city"`
	Temp_C float32 `json:"temp_C"`
	Temp_F float32 `json:"temp_F"`
	Temp_K float32 `json:"temp_K"`
}

func NewMs2() *Ms2 {
	return &Ms2{}
}

func (b *Ms2) FindData(ctx context.Context, zipcode *entity.ZipCode) (*entity.Temperature, error) {

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/temperature?zipcode="+zipcode.Cep, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Ms2
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &data.Success)
		if err != nil {
			log.Println(err)
			return nil, errors.New("internal error")
		}

		return &entity.Temperature{City: data.Success.Message.City, Temp_C: data.Success.Message.Temp_C, Temp_F: data.Success.Message.Temp_F, Temp_K: data.Success.Message.Temp_K}, nil

	}

	err = json.Unmarshal(body, &data.Error)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	return &entity.Temperature{}, errors.New(data.Error.Message)

}
