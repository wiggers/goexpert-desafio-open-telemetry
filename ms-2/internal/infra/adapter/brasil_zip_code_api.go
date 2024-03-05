package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	lib "github.com/wiggers/goexpert/desafio/1-temperatura/pkg"
)

type BrasilApiData struct {
	City string `json:"city"`
}

func NewBrasilApiData() *BrasilApiData {
	return &BrasilApiData{}
}

func (b *BrasilApiData) FindCity(ctx context.Context, zipcode *entity.ZipCode) (*entity.City, error) {

	dataCep, err := lib.CallHttpGet(ctx, "https://brasilapi.com.br/api/cep/v1/"+zipcode.Cep)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	var data BrasilApiData
	err = json.Unmarshal(dataCep, &data)

	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	return &entity.City{City: data.City}, nil
}
