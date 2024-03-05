package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAValidZipCode(t *testing.T) {
	zipcode, err := NewZipCode("01153000")
	assert.Nil(t, err)
	assert.Equal(t, "01153000", zipcode.Cep)

}

func TestGivenAInvalidZipCode(t *testing.T) {
	_, err := NewZipCode("01153-000")
	assert.Error(t, err, "invalid zip code")
}
