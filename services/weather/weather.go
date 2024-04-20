package weather

import "github.com/woodjeffrey2/weather-service/models"

// GetCurrentWeather returns the current weather for the provided coordinates
func (s *service) GetCurrentWeather(lat, lon float64) (models.CurrentWeather, error) {
	// TODO: Implement
	weather := models.CurrentWeather{}
	return weather, nil
}

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
