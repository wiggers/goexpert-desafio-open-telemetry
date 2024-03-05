package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAWeather_ThenShouldCalculateKelvinFahrenheit(t *testing.T) {
	var weather Weather
	weather.Temp_C = 29
	assert.Equal(t, float32(84.2), weather.GetFahrenheit())
	assert.Equal(t, float32(302), weather.GetKelvin())
}
