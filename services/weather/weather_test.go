package weather

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woodjeffrey2/weather-service/models"
)

func TestDescribeTemp(t *testing.T) {
	var myTests = map[string]struct {
		temp         float64
		expectedDesc string
	}{
		"Given a temp higher than 80 expect to return 'hot'": {
			temp:         100,
			expectedDesc: "hot",
		},
		"Given a temp between 50 and 80 expect to return 'moderate'": {
			temp:         65,
			expectedDesc: "moderate",
		},
		"Given a temp below 50 expect to return 'cold'": {
			temp:         25,
			expectedDesc: "cold",
		},
	}
	for _, tc := range myTests {
		assert.Equal(t, tc.expectedDesc, describeTemp(tc.temp))
	}
}

func TestGetCurrentWeather(t *testing.T) {
	var myTests = map[string]struct {
		lat          float64
		lon          float64
		mockRespBody []byte
		mockStatus   int
		expectedResp models.CurrentWeather
		expectedErr  string
	}{
		"Given api call success expect to return correct summary": {
			lat: 1.53,
			lon: 23.46,
			mockRespBody: []byte(`
			{
				"weather": [
					{
						"description": "moderate rain"
					}
				],
				"main": {
					"temp": 57.3
				}
			}
		`),
			mockStatus: http.StatusOK,
			expectedResp: models.CurrentWeather{
				Latitude:        1.53,
				Longitude:       23.46,
				Condition:       "moderate rain",
				TempDescription: "moderate",
			},
		},
		"Given error fetching weather expect to return error": {
			lat:          1.53,
			lon:          23.46,
			mockRespBody: []byte(`{"error":"something went wrong"}`),
			mockStatus:   http.StatusInternalServerError,
			expectedErr:  "fetching weather report: openweather api returned an invalid response. Status code: 500 Response: {\"error\":\"something went wrong\"}",
		},
		"Given valid response with no weather object expect to return error": {
			lat: 1.53,
			lon: 23.46,
			mockRespBody: []byte(`
			{
				"weather": [],
				"main": {
					"temp": 57.3
				}
			}
		`),
			mockStatus:  http.StatusOK,
			expectedErr: "no weather returned for lat: 1.530000 lon: 23.460000",
		},
	}
	for _, tc := range myTests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tc.mockStatus)
			w.Write(tc.mockRespBody)
		}))
		defer server.Close()

		s := service{
			log:       slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			client:    &http.Client{},
			owBaseUrl: server.URL,
		}
		resp, err := s.GetCurrentWeather(tc.lat, tc.lon)
		if tc.expectedErr != "" {
			assert.EqualError(t, err, tc.expectedErr)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResp, resp)
		}
	}
}
