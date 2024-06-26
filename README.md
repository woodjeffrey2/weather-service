# weather-service
Simple Golang API service for fetching a summary of the current weather for a given latitude and longitude

- [weather-service](#weather-service)
  - [Requirements](#requirements)
    - [Running the Server](#running-the-server)
    - [Development](#development)
  - [Local Setup](#local-setup)
  - [Commands](#commands)
  - [Testing](#testing)
  - [Design](#design)
    - [API](#api)
      - [Get Weather Report](#get-weather-report)
    - [Service Architecture](#service-architecture)
    - [Next Steps](#next-steps)

## Requirements
### Running the Server
* Golang >= 1.22.2 or Docker / Rancher Desktop
  * [https://go.dev/doc/install](https://go.dev/doc/install)
  * [https://docs.docker.com/desktop/install/mac-install/](https://docs.docker.com/desktop/install/mac-install/)
* OpenWeather API Key -
  * [https://openweathermap.org/faq](https://openweathermap.org/faq)

### Development
* Mockery >= 2.42.3 -
  * [https://vektra.github.io/mockery/latest/installation/](https://vektra.github.io/mockery/latest/installation/)

## Local Setup
Set the OpenWeather API key env var:
```sh
% export OW_API_KEY=<your API key>
```
Run the API server locally:
```sh
% go mod download
% make run
```
Or spin up a local Docker container:
```sh
% make run-docker
```
Make a request to the local `GET /weather-report` endpoint:
```sh
% curl -G -d 'lat=12.34' -d 'lon=56.78' http://localhost:8080/weather-report
```
## Commands
| Command           | Description              |
| ----------------- | ------------------------ |
| `make test`       | Run unit tests           |
| `make build`      | Compile binary           |
| `make run`        | Run server locally       |
| `make run-docker` | Run server with Docker   |
| `make mocks`      | Generate interface mocks |

## Testing
Unit tests can be run with the `make test` command locally.

There is also a `Github Actions` workflow defined for the service that runs the tests on any code push. All unit tests must pass before PRs can be merged to `main`.

## Design
This service was created as a coding exercise and is an MVP for an API server with a single endpoint.

### API
#### Get Weather Report
Calls the [OpenWeather API](https://openweathermap.org/current) and returns a summary report of the weather conditions and temperature for a given location.

Example - `GET http://localhost:8080/weather-report?lat=12.34&lon=56.78`

Query Params
* `lat` - latitude of the location for the weather report (Default: 0)
* `lon` - longitude of the location for the weather report (Default: 0)

Example Response (200 status)
```json
{
  "data":{
    "latitude": 12.34,
    "longitude": 56.78,
    "condition": "scattered clouds",
    "temp_description": "hot"
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

Since the Services layer has no dependencies on the Golang http server implementation, the interface, implementation, and business logic could be reused with a completely different service architecture without requiring any modifications.

### Next Steps
We're starting off with a simple MVP that can be compiled and run locally for the purposes of the coding exercise.

There are some improvements and modifications that would be at the top of my list to build out on a production service:
* Add unit tests to `Github Actions` CI
* Add test coverage and PR thresholds in `Gihub Actions` CI
* Add APM and automated alerting
* Create a wrapper / Handler implementation for the structured logger
* Add an `AWS Lambda` + `API Gateway` API implementation using `CDK` for IaC
* Automate deployments to AWS cloud with `Github Actions`
* Add ephemeral stack creation in `Github Actions` for integration testing changes in the AWS cloud
