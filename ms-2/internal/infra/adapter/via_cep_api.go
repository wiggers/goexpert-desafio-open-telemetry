package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	lib "github.com/wiggers/goexpert/desafio/1-temperatura/pkg"
)

type ViaCepApiData struct {
	City string `json:"localidade"`
}

func NewViaCepApiData() *ViaCepApiData {
	return &ViaCepApiData{}
}

func (b *ViaCepApiData) FindCity(ctx context.Context, zipcode *entity.ZipCode) (*entity.City, error) {

	dataCep, err := lib.CallHttpGet(ctx, "https://viacep.com.br/ws/"+zipcode.Cep+"/json/")
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	var data ViaCepApiData
	err = json.Unmarshal(dataCep, &data)

	if err != nil {
		log.Println(err)
		return nil, errors.New("internal error")
	}

	return &entity.City{City: data.City}, nil
}
