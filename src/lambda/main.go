package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/woodjeffrey2/weather-service/handlers/lambdahandlers"
	"github.com/woodjeffrey2/weather-service/handlers/lambdahandlers/report"
	"github.com/woodjeffrey2/weather-service/services/weather"
)

const (
	HTTP_PORT      = ":8080"
	OW_BASE_URL    = "https://api.openweathermap.org"
	CLIENT_TIMEOUT = 5
)

var (
	weatherHandler lambdahandlers.WeatherReportHandler
	logger         *slog.Logger
)

func init() {
	// inject dependencies and initialize the handler
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	client := &http.Client{
		Timeout: CLIENT_TIMEOUT * time.Second,
	}
	service := weather.NewService(logger, client, OW_BASE_URL)
	weatherHandler = report.NewHandler(logger, service)
}

func main() {
	lambda.Start(weatherHandler)
}
