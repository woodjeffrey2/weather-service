package weather

import (
	"errors"
	"fmt"

	"github.com/woodjeffrey2/weather-service/models"
)

var (
	ErrNoReport = errors.New("no weather report")
)

// describeTemp returns a string description of "hot", "moderate", or "cold"
// based on the input temperature (degrees Farenheit)
func describeTemp(temp float64) string {
	switch {
	case temp > 80:
		return "hot"
	case temp > 50:
		return "moderate"
	default:
		return "cold"
	}
}

// GetCurrentWeather returns the current weather conditions for the provided coordinates
func (s *service) GetCurrentWeather(lat, lon float64) (models.CurrentWeather, error) {
	report, err := s.fetchOWCurrent(lat, lon)
	if err != nil {
		return models.CurrentWeather{}, fmt.Errorf("fetching weather report: %w", err)
	}
	if len(report.Weather) < 1 {
		return models.CurrentWeather{}, fmt.Errorf("no weather returned for lat: %f lon: %f", lat, lon)
	}

	return models.CurrentWeather{
		Latitude:        lat,
		Longitude:       lon,
		Condition:       report.Weather[0].Description,
		TempDescription: describeTemp(report.Main.Temp),
	}, nil
}
