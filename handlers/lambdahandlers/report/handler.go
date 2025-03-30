package report

import (
	"log/slog"

	"github.com/woodjeffrey2/weather-service/handlers/lambdahandlers"
	"github.com/woodjeffrey2/weather-service/services"
)

// reportLambdaHandler implements WeatherReportHandler
type reportLambdaHandler struct {
	log     *slog.Logger
	weather services.WeatherService
}

// make sure the interface is implemented by reportHandler
var _ lambdahandlers.WeatherReportHandler = new(reportLambdaHandler)

// NewHandler instantiates a new WeatherReportHandler implementation
func NewHandler(l *slog.Logger, s services.WeatherService) lambdahandlers.WeatherReportHandler {
	return &reportLambdaHandler{
		log:     l,
		weather: s,
	}
}
