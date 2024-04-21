package report

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mocks "github.com/woodjeffrey2/weather-service/mocks/github.com/woodjeffrey2/weather-service/services"
	"github.com/woodjeffrey2/weather-service/models"
)

func TestParseLatLon(t *testing.T) {
	var myTests = map[string]struct {
		params      url.Values
		expectedLat float64
		expectedLon float64
		expectedErr string
	}{
		"Given params present expect to return correct values": {
			params: url.Values{
				"lat": []string{"23.45"},
				"lon": []string{"67.89"},
			},
			expectedLat: 23.45,
			expectedLon: 67.89,
		},
		"Given missing params expect to default to 0": {
			params:      url.Values{},
			expectedLat: 0,
			expectedLon: 0,
		},
		"Given invalid value expect to return error": {
			params: url.Values{
				"lat": []string{"Boston"},
			},
			expectedErr: "parsing lat query param: strconv.ParseFloat: parsing \"Boston\": invalid syntax",
		},
	}
	for _, tc := range myTests {
		lat, lon, err := parseLatLon(tc.params)
		if tc.expectedErr != "" {
			assert.EqualError(t, err, tc.expectedErr)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tc.expectedLat, lat)
			assert.Equal(t, tc.expectedLon, lon)
		}
	}
}

func TestWeatherReportHandler(t *testing.T) {
	var myTests = map[string]struct {
		latStr         string
		lonStr         string
		getWeatherResp models.CurrentWeather
		getWeatherErr  error
		paramErr       bool
		expectedResp   string
		expectedStatus int
	}{
		"Given valid request expect to return correct response": {
			latStr: "12.34",
			lonStr: "56.78",
			getWeatherResp: models.CurrentWeather{
				Latitude:        12.34,
				Longitude:       56.78,
				Condition:       "slightly cloudy",
				TempDescription: "cold",
			},
			expectedStatus: http.StatusOK,
			expectedResp:   `{"data":{"Latitude":12.34,"Longitude":56.78,"Condition":"slightly cloudy","TempDescription":"cold"}}`,
		},
		"Given OpenWeather API error expect to return 500 response": {
			latStr: "12.34",
			lonStr: "56.78",
			getWeatherResp: models.CurrentWeather{
				Latitude:  12.34,
				Longitude: 56.78,
			},
			getWeatherErr:  errors.New("API error"),
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   `API error`,
		},
		"Given invalid query params expect to return 400 response": {
			latStr:         "notanumber",
			lonStr:         "56.78",
			paramErr:       true,
			expectedStatus: http.StatusBadRequest,
			expectedResp:   `parsing lat query param: strconv.ParseFloat: parsing "notanumber": invalid syntax`,
		},
	}
	for _, tc := range myTests {
		mService := mocks.NewMockWeatherService(t)
		if !tc.paramErr {
			mService.On("GetCurrentWeather", tc.getWeatherResp.Latitude, tc.getWeatherResp.Longitude).
				Return(tc.getWeatherResp, tc.getWeatherErr)
		}
		weatherHandler := reportHandler{
			log:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			weather: mService,
		}

		req, err := http.NewRequest("GET", "/weather-service", nil)
		require.NoError(t, err)
		q := req.URL.Query()
		q.Add(LAT_PARAM, tc.latStr)
		q.Add(LON_PARAM, tc.lonStr)
		req.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(weatherHandler.WeatherReportHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedStatus, rr.Code)
		assert.Equal(t, tc.expectedResp+"\n", rr.Body.String())
	}
}
