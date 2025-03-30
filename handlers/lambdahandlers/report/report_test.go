package report

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/woodjeffrey2/weather-service/handlers/presenters"
	mocks "github.com/woodjeffrey2/weather-service/mocks/github.com/woodjeffrey2/weather-service/services"
	"github.com/woodjeffrey2/weather-service/models"
)

func TestWeatherReportHandler(t *testing.T) {
	var myTests = map[string]struct {
		req            events.APIGatewayProxyRequest
		getWeatherResp models.CurrentWeather
		getWeatherErr  error
		paramErr       bool
		expectedResp   events.APIGatewayProxyResponse
	}{
		"Given valid request expect to return correct response": {
			req: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					presenters.LAT_PARAM: "12.34",
					presenters.LON_PARAM: "56.78",
				},
			},
			getWeatherResp: models.CurrentWeather{
				Latitude:        12.34,
				Longitude:       56.78,
				Condition:       "slightly cloudy",
				TempDescription: "cold",
			},
			expectedResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       `{"data":{"latitude":12.34,"longitude":56.78,"condition":"slightly cloudy","temp_description":"cold"}}`,
			},
		},
		"Given OpenWeather API error expect to return 500 response": {
			req: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					presenters.LAT_PARAM: "12.34",
					presenters.LON_PARAM: "56.78",
				},
			},
			getWeatherResp: models.CurrentWeather{
				Latitude:  12.34,
				Longitude: 56.78,
			},
			getWeatherErr: errors.New("API error"),
			expectedResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `API error`,
			},
		},
		"Given invalid query params expect to return 400 response": {
			req: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					presenters.LAT_PARAM: "notanumber",
					presenters.LON_PARAM: "56.78",
				},
			},
			paramErr: true,
			expectedResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `strconv.ParseFloat: parsing "notanumber": invalid syntax`,
			},
		},
	}
	for _, tc := range myTests {
		mService := mocks.NewMockWeatherService(t)
		if !tc.paramErr {
			mService.On("GetCurrentWeather", tc.getWeatherResp.Latitude, tc.getWeatherResp.Longitude).
				Return(tc.getWeatherResp, tc.getWeatherErr)
		}
		weatherHandler := reportLambdaHandler{
			log:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			weather: mService,
		}

		response, err := weatherHandler.WeatherReportHandler(context.Background(), tc.req)

		assert.NoError(t, err)
		assert.Equal(t, tc.expectedResp, response)
	}
}
