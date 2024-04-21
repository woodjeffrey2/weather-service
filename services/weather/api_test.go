package weather

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchOWCurrent(t *testing.T) {
	var myTests = map[string]struct {
		lat          float64
		lon          float64
		mockRespBody []byte
		mockStatus   int
		expectedResp owCurrentResponse
		expectedErr  string
	}{
		"Given 200 response expect to return correct weather info": {
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
					"temp": 298.48
				}
			}
		`),
			mockStatus: http.StatusOK,
			expectedResp: owCurrentResponse{
				Weather: []owWeather{
					{
						Description: "moderate rain",
					},
				},
				Main: owMain{
					Temp: 298.48,
				},
			},
		},
		"Given invalid response from the API expect to return error": {
			lat:          1.53,
			lon:          23.46,
			mockRespBody: []byte(`{"error": "something went wrong"}`),
			mockStatus:   http.StatusInternalServerError,
			expectedErr:  "openweather api returned an invalid response. Status code: 500 Response: {\"error\": \"something went wrong\"}",
		},
	}
	for _, tc := range myTests {
		assert.NoError(t, os.Setenv("OW_API_KEY", "1233456"))
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/data/2.5/weather", r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("Accept"))
			assert.Equal(t, "appid=1233456&lat=1.53&lon=23.46&units=imperial", r.URL.RawQuery)

			w.WriteHeader(tc.mockStatus)
			w.Write(tc.mockRespBody)
		}))
		defer server.Close()

		s := service{
			client:    &http.Client{},
			owBaseUrl: server.URL,
		}
		resp, err := s.fetchOWCurrent(tc.lat, tc.lon)
		if tc.expectedErr != "" {
			assert.EqualError(t, err, tc.expectedErr)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResp, resp)
		}
	}
}
