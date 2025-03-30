package report

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/woodjeffrey2/weather-service/handlers/presenters"
)

// WeatherReportHandler handles requests to the /weather-report path
func (h *reportLambdaHandler) WeatherReportHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.log.Info("weather report request",
		"method", request.HTTPMethod,
		"path", request.Path,
		"query", request.QueryStringParameters,
	)

	// parse the query params to get latitude & longitude
	lat, lon, err := parseLatLon(request.QueryStringParameters)
	if err != nil {
		return h.handleError(err, 400)
	}

	// call weather service to get current weather
	weather, err := h.weather.GetCurrentWeather(lat, lon)
	if err != nil {
		return h.handleError(err, 500)
	}

	responseBody, err := json.Marshal(presenters.WeatherReportResponse{Data: presenters.FromModel(weather)})
	if err != nil {
		return h.handleError(err, 500)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

// handleError logs the error and returns an APIGatewayProxyResponse with the error message
func (h *reportLambdaHandler) handleError(err error, statusCode int) (events.APIGatewayProxyResponse, error) {
	h.log.Error("error", "error", err)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       err.Error(),
	}, nil
}

func parseLatLon(params map[string]string) (float64, float64, error) {
	latStr := params[presenters.LAT_PARAM]
	lonStr := params[presenters.LON_PARAM]
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return 0, 0, err
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return 0, 0, err
	}
	return lat, lon, nil
}
