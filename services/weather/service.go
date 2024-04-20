package weather

import (
	"net/http"

	"github.com/woodjeffrey2/weather-service/services"
)

// service implements the WeatherService interface
type service struct {
	client    *http.Client
	owBaseUrl string
}

// NewService instantiates a new WeatherService implementation
func NewService(cli *http.Client, baseURL string) services.WeatherService {
	return &service{
		client:    cli,
		owBaseUrl: DEFAULT_BASE_URL,
	}
}
