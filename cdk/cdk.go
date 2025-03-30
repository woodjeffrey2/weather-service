package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create the weather report lambda
	getWeatherReportLambda := awslambda.NewFunction(stack, jsii.String("WeatherReportLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("main"),
		Code:    awslambda.Code_FromAsset(jsii.String("./src/lambda"), nil),
		Environment: &map[string]*string{
			"OW_BASE_URL": jsii.String("https://api.openweathermap.org"),
		},
	})

	// create the http aws api gateway
	httpApi := awsapigatewayv2.NewHttpApi(stack, jsii.String("WeatherApi"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String("WeatherApi"),
	})

	// create the integration for the weather report lambda
	weatherIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		jsii.String("GetWeatherReport"),
		getWeatherReportLambda,
		&awsapigatewayv2integrations.HttpLambdaIntegrationProps{},
	)

	// add the route for the weather report lambda
	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/weather-report"),
		Integration: weatherIntegration,
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_GET,
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
