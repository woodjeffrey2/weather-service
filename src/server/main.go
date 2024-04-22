package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/woodjeffrey2/weather-service/handlers"
	"github.com/woodjeffrey2/weather-service/handlers/report"
	"github.com/woodjeffrey2/weather-service/services/weather"
)

const (
	HTTP_PORT      = ":8080"
	OW_BASE_URL    = "https://api.openweathermap.org"
	CLIENT_TIMEOUT = 5
)

var (
	weatherHandler handlers.WeatherReportHandler
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
	http.HandleFunc("/weather-report", weatherHandler.WeatherReportHandler)

	logger.Info("Server is running.", "URL", fmt.Sprintf("http://localhost%s", HTTP_PORT))
	log.Fatal(http.ListenAndServe(HTTP_PORT, nil))
}
