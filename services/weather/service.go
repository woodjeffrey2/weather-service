package weather

import (
	"log/slog"
	"net/http"

	"github.com/woodjeffrey2/weather-service/services"
)

// service implements the WeatherService interface
type service struct {
	log       *slog.Logger
	client    *http.Client
	owBaseUrl string
}

// NewService instantiates a new WeatherService implementation
func NewService(log *slog.Logger, cli *http.Client, baseURL string) services.WeatherService {
	return &service{
		log:       log,
		client:    cli,
		owBaseUrl: baseURL,
	}
}
