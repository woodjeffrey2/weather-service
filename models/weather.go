package models

// CurrentWeather DTO for current weather data
type CurrentWeather struct {
	Latitude        float64
	Longitude       float64
	Condition       string
	TempDescription string
}
