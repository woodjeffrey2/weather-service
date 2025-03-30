package presenters

import "github.com/woodjeffrey2/weather-service/models"

const (
	LAT_PARAM = "lat"
	LON_PARAM = "lon"
)

// WeatherReportResponse JSON response for GET
type WeatherReportResponse struct {
	Data CurrentWeather `json:"data"`
}

// CurrentWeather JSON response for current weather summary
type CurrentWeather struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Condition       string  `json:"condition"`
	TempDescription string  `json:"temp_description"`
}

// FromModel converts a CurrentWeather model to the currentWeather presentataion struct
func FromModel(model models.CurrentWeather) CurrentWeather {
	return CurrentWeather(model)
}
