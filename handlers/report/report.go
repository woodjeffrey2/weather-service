package report

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/woodjeffrey2/weather-service/handlers"
	"github.com/woodjeffrey2/weather-service/models"
	"github.com/woodjeffrey2/weather-service/services"
)

const (
	LAT_PARAM = "lat"
	LON_PARAM = "lon"
)

// reportHandler implements WeatherReportHandler
type reportHandler struct {
	weather services.WeatherService
}

// NewWeatherReportHandler instantiates a new WeatherReportHandler implementation
func NewWeatherReportHandler(s services.WeatherService) handlers.WeatherReportHandler {
	return &reportHandler{weather: s}
}

// weatherReportResponse JSON response for GET /weather-report
type weatherReportResponse struct {
	Data models.CurrentWeather `json:"data"`
	// Latitude  float64 `json:"latitude"`
	// Longitude float64 `json:"longitude"`
	// Condition string  `json:"conditions"`
	// TempRange string  `json:"temperature"`
}

// WeatherReportHandler handles requests to the /weather-report path
func (h *reportHandler) WeatherReportHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getWeatherReport(w, r)
	default:
		http.Error(w, "Method %s not allowed", http.StatusMethodNotAllowed)
	}
}

// getWeatherReport returns a JSON response with a summary of the current weather
func (h *reportHandler) getWeatherReport(w http.ResponseWriter, r *http.Request) {
	// parse the query params to get lat & lon
	lat, lon, err := parseLatLon(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weather, err := h.weather.GetCurrentWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := weatherReportResponse{Data: weather}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// parseLatLon parses the lat and lon from request query params
func parseLatLon(params url.Values) (lat float64, lon float64, err error) {
	if lp, ok := params[LAT_PARAM]; ok {
		lat, err = strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return lat, lon, fmt.Errorf("parsing lat query param: %w", err)
		}
	}
	if lp, ok := params[LON_PARAM]; ok {
		lon, err = strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return lat, lon, fmt.Errorf("parsing lon query param: %w", err)
		}
	}
	return
}
