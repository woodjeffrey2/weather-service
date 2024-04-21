package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/woodjeffrey2/weather-service/handlers"
	"github.com/woodjeffrey2/weather-service/handlers/report"
	"github.com/woodjeffrey2/weather-service/services/weather"
)

const (
	HTTP_PORT   = ":8080"
	OW_BASE_URL = "https://api.openweathermap.org"
)

var (
	weatherHandler handlers.WeatherReportHandler
)

func init() {
	service := weather.NewService(&http.Client{}, OW_BASE_URL)
	weatherHandler = report.NewWeatherReportHandler(service)
}

func main() {
	// service := weather.NewService(&http.Client{}, OW_BASE_URL)
	// handler := report.NewWeatherReportHandler(service)
	http.HandleFunc("/weather-report", weatherHandler.WeatherReportHandler)

	fmt.Printf("Server is running at http://localhost%s", HTTP_PORT)
	log.Fatal(http.ListenAndServe(HTTP_PORT, nil))
}
