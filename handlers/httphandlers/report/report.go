package report

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/woodjeffrey2/weather-service/handlers/presenters"
)

// WeatherReportHandler handles requests to the /weather-report path
func (h *reportHandler) WeatherReportHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Info("weather report request",
		"method", r.Method,
		"path", r.URL.Path,
		"query", r.URL.RawQuery,
	)
	switch r.Method {
	case "GET":
		h.getWeatherReport(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
	}
}

// getWeatherReport returns a JSON response with a summary of the current weather
func (h *reportHandler) getWeatherReport(w http.ResponseWriter, r *http.Request) {
	// parse the query params to get lat & lon
	lat, lon, err := parseLatLon(r.URL.Query())
	if err != nil {
		h.log.Error("error parsing query params", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("get weather report", "lat", lat, "lon", lon)
	weather, err := h.weather.GetCurrentWeather(lat, lon)
	if err != nil {
		h.log.Error("error getting current weather", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := presenters.WeatherReportResponse{Data: presenters.FromModel(weather)}
	h.log.Info("report generated", "response", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// parseLatLon parses the lat and lon from request query params
func parseLatLon(params url.Values) (lat float64, lon float64, err error) {
	if lp, ok := params[presenters.LAT_PARAM]; ok && len(lp) > 0 {
		lat, err = strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return lat, lon, fmt.Errorf("parsing lat query param: %w", err)
		}
	}
	if lp, ok := params[presenters.LON_PARAM]; ok && len(lp) > 0 {
		lon, err = strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return lat, lon, fmt.Errorf("parsing lon query param: %w", err)
		}
	}
	return
}
