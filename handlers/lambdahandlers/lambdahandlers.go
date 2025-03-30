package lambdahandlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type WeatherReportHandler interface {
	WeatherReportHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
