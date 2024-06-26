// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockWeatherReportHandler is an autogenerated mock type for the WeatherReportHandler type
type MockWeatherReportHandler struct {
	mock.Mock
}

type MockWeatherReportHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWeatherReportHandler) EXPECT() *MockWeatherReportHandler_Expecter {
	return &MockWeatherReportHandler_Expecter{mock: &_m.Mock}
}

// WeatherReportHandler provides a mock function with given fields: w, r
func (_m *MockWeatherReportHandler) WeatherReportHandler(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockWeatherReportHandler_WeatherReportHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WeatherReportHandler'
type MockWeatherReportHandler_WeatherReportHandler_Call struct {
	*mock.Call
}

// WeatherReportHandler is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockWeatherReportHandler_Expecter) WeatherReportHandler(w interface{}, r interface{}) *MockWeatherReportHandler_WeatherReportHandler_Call {
	return &MockWeatherReportHandler_WeatherReportHandler_Call{Call: _e.mock.On("WeatherReportHandler", w, r)}
}

func (_c *MockWeatherReportHandler_WeatherReportHandler_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockWeatherReportHandler_WeatherReportHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockWeatherReportHandler_WeatherReportHandler_Call) Return() *MockWeatherReportHandler_WeatherReportHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWeatherReportHandler_WeatherReportHandler_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockWeatherReportHandler_WeatherReportHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWeatherReportHandler creates a new instance of MockWeatherReportHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWeatherReportHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWeatherReportHandler {
	mock := &MockWeatherReportHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
