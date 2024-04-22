package report

import (
	"log/slog"

	"github.com/woodjeffrey2/weather-service/handlers"
	"github.com/woodjeffrey2/weather-service/services"
)

// reportHandler implements WeatherReportHandler
type reportHandler struct {
	log     *slog.Logger
	weather services.WeatherService
}

// make sure the interface is implemented by reportHandler
var _ handlers.WeatherReportHandler = new(reportHandler)

// NewHandler instantiates a new WeatherReportHandler implementation
func NewHandler(l *slog.Logger, s services.WeatherService) handlers.WeatherReportHandler {
	return &reportHandler{
		log:     l,
		weather: s,
	}
}
