package weather

import (
	"log/slog"
	"net/http"

	"github.com/woodjeffrey2/weather-service/services"
)

// weatherService implements the WeatherService interface
type weatherService struct {
	log       *slog.Logger
	client    *http.Client
	owBaseUrl string
}

// make sure the interface is implemented by reportHandler
var _ services.WeatherService = new(weatherService)

// NewService instantiates a new WeatherService implementation
func NewService(log *slog.Logger, cli *http.Client, baseURL string) services.WeatherService {
	return &weatherService{
		log:       log,
		client:    cli,
		owBaseUrl: baseURL,
	}
}
