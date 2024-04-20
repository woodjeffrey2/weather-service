package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/woodjeffrey2/weather-service/models"
	"github.com/woodjeffrey2/weather-service/services"
)

const (
	DEFAULT_BASE_URL = "https://api.openweathermap.org"
	OW_CURRENT_PATH  = "/data/2.5/weather"
	OW_API_KEY_VAR   = "OW_API_KEY"
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

// GetCurrentWeather returns the current weather for the provided coordinates
func (s *service) GetCurrentWeather(lat, lon float64) (models.CurrentWeather, error) {
	// TODO: Implement
	weather := models.CurrentWeather{}
	return weather, nil
}

// owCurrentResponse API response struct for OpenWeather current endpoint
// https://openweathermap.org/current
type owCurrentResponse struct {
	Weather []owWeather `json:"weather"`
	Main    owMain      `json:"main"`
}
type owWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}
type owMain struct {
	Temp float64 `json:"temp"`
}

// fetchOWCurrent fetches the current weather from the OpenWeather API
func (w *service) fetchOWCurrent(lat, lon float64) (owCurrentResponse, error) {
	var owCurrent owCurrentResponse

	path, err := url.JoinPath(w.owBaseUrl, OW_CURRENT_PATH)
	if err != nil {
		return owCurrent, fmt.Errorf("constructing API path: %w", err)
	}

	// create http request
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return owCurrent, fmt.Errorf("creating http request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	// set query params
	q := req.URL.Query()
	q.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	q.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))
	q.Add("appid", os.Getenv(OW_API_KEY_VAR))
	req.URL.RawQuery = q.Encode()

	// execute http request
	resp, err := w.client.Do(req)
	if err != nil {
		return owCurrent, fmt.Errorf("executing http request: %w", err)
	}

	// unmarshal API response body to struct
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&owCurrent)
	if err != nil {
		return owCurrent, fmt.Errorf("decoding response: %w", err)
	}
	return owCurrent, nil
}
