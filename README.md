# weather-service
Simple Golang API service for fetching a summary of the current weather for a given latitude and longitude

## Requirements
* Golang >= 1.22.2 - [https://go.dev/doc/install](https://go.dev/doc/install)
* Mockery >= 2.42.3 - [https://vektra.github.io/mockery/latest/installation/](https://vektra.github.io/mockery/latest/installation/)

## Setup
Set the OpenWeather API key env var:
```sh
% export OW_API_KEY=<your API key>
```
Start the local API server:
```sh
% make run
```
Make a request to the local `GET /weather-report` endpoint:
```sh
% curl -G -d 'lat=12.34' -d 'lon=56.78' http://localhost:8080/weather-report
```
## Commands
| Command       | Description              |
| ------------- | ------------------------ |
| `make test`   | Run unit tests           |
| `make build`  | Compile binary           |
| `make run`    | Run server locally       |
| `make mocks`  | Generate interface mocks |

## Design
This service was created as a coding exercise and is a MVP for an API server with 1 endpoint.

### API
#### Get Weather Report
Returns a summary report of the weather conditions and temperature for a given location.

Example - `GET http://localhost:8080/weather-report?lat=12.34&lon=56.78`

Query Params
* `lat` - latitude of the location for the weather report (Default: 0)
* `lon` - longitude of the location for the weather report (Default: 0)

Example Response
```json
{
  "data":{
    "Latitude":12.34,
    "Longitude":56.78,
    "Condition":"scattered clouds",
    "TempDescription":"hot"
  }
}
```
### Service Architecture
For this service we're using a layered  architecture based on Uncle Bob's [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) pattern.

![Architecture Diagram](/docs/images/weather-service-layers.png)

In order to reduce coupling between layers, provide isolation for unit tests, and increase maintainability, we rely on dependency inversion and Golang interfaces for communication between layers.

![Architecture Diagram](/docs/images/weather-service-deps.png)

Mocks for each interface are generated using the `mockery` library and ensure that we can write unit tests for each layer in isolation.

This loose coupling also makes it much simpler to make changes like swapping the entire Golang http server implementation with an implementation that uses API Gateway + Lambda to serve http requests instead.

Since the Services layer has no dependencies on the Golang http server implementation, that interface and service implementation could be reused for a completely different server architecture without requiring any modifications.

### Next Steps
We're starting off with a simple MVP that can be compiled and run locally for the purposes of the coding exercise.

There are some improvements and modifications that would be at the top of my list to build out on a production service:
* Add `Docker` and `Docker Compose` setup for running the http server locally
* Run unit tests in `Github Actions` CI
* Add test coverage calculation and thresholds in Gihub Actions CI
* Add an AWS Lambda + API Gateway API implementation using CDK
* Automate deployments to AWS cloud with CDK Lambda + API Gateway implementation
* Add APM and automated alerting
* Create a wrapper / Handler implementation for the structured logger
* Add ephemeral stack creation in Github Actions CI for integration testing changes in the AWS cloud
