package weather

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindEmails(t *testing.T) {
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
				"coord": {
					"lon": 1.53,
					"lat": 23.46
				},
				"weather": [
					{
						"id": 501,
						"main": "Rain",
						"description": "moderate rain"
					}
				],
				"base": "stations",
				"main": {
					"temp": 298.48
				}
			}
		`),
			mockStatus: http.StatusOK,
			expectedResp: owCurrentResponse{
				Weather: []owWeather{
					{
						Main:        "Rain",
						Description: "moderate rain",
					},
				},
				Main: owMain{
					Temp: 298.48,
				},
			},
		},
	}
	for _, tc := range myTests {
		assert.NoError(t, os.Setenv("OW_API_KEY", "1233456"))
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/data/2.5/weather", r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("Accept"))
			assert.Equal(t, "appid=1233456&lat=1.53&lon=23.46", r.URL.RawQuery)

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
