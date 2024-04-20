package services

import (
	"github.com/woodjeffrey2/weather-service/models"
)

// WeatherService interface for interacting with weather data
type WeatherService interface {
	GetCurrentWeather(lat, lon float64) (models.CurrentWeather, error)
}
