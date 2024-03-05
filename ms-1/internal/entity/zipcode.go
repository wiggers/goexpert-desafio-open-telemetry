package entity

import (
	"errors"
	"regexp"
)

type ZipCode struct {
	Cep string
}

func NewZipCode(cep string) (ZipCode, error) {
	valid := validCep(cep)
	if valid {
		return ZipCode{Cep: cep}, nil
	}

	return ZipCode{}, errors.New("invalid zip code")

}

func validCep(cep string) bool {
	return regexp.MustCompile(`^\d{8}$`).MatchString(cep)
}
