// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/woodjeffrey2/weather-service/models"
)

// MockWeatherService is an autogenerated mock type for the WeatherService type
type MockWeatherService struct {
	mock.Mock
}

type MockWeatherService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWeatherService) EXPECT() *MockWeatherService_Expecter {
	return &MockWeatherService_Expecter{mock: &_m.Mock}
}

// GetCurrentWeather provides a mock function with given fields: lat, lon
func (_m *MockWeatherService) GetCurrentWeather(lat float64, lon float64) (models.CurrentWeather, error) {
	ret := _m.Called(lat, lon)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentWeather")
	}

	var r0 models.CurrentWeather
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, float64) (models.CurrentWeather, error)); ok {
		return rf(lat, lon)
	}
	if rf, ok := ret.Get(0).(func(float64, float64) models.CurrentWeather); ok {
		r0 = rf(lat, lon)
	} else {
		r0 = ret.Get(0).(models.CurrentWeather)
	}

	if rf, ok := ret.Get(1).(func(float64, float64) error); ok {
		r1 = rf(lat, lon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWeatherService_GetCurrentWeather_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentWeather'
type MockWeatherService_GetCurrentWeather_Call struct {
	*mock.Call
}

// GetCurrentWeather is a helper method to define mock.On call
//   - lat float64
//   - lon float64
func (_e *MockWeatherService_Expecter) GetCurrentWeather(lat interface{}, lon interface{}) *MockWeatherService_GetCurrentWeather_Call {
	return &MockWeatherService_GetCurrentWeather_Call{Call: _e.mock.On("GetCurrentWeather", lat, lon)}
}

func (_c *MockWeatherService_GetCurrentWeather_Call) Run(run func(lat float64, lon float64)) *MockWeatherService_GetCurrentWeather_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64))
	})
	return _c
}

func (_c *MockWeatherService_GetCurrentWeather_Call) Return(_a0 models.CurrentWeather, _a1 error) *MockWeatherService_GetCurrentWeather_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWeatherService_GetCurrentWeather_Call) RunAndReturn(run func(float64, float64) (models.CurrentWeather, error)) *MockWeatherService_GetCurrentWeather_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWeatherService creates a new instance of MockWeatherService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWeatherService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWeatherService {
	mock := &MockWeatherService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}