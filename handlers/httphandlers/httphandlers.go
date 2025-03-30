package httphandlers

import "net/http"

type WeatherReportHandler interface {
	WeatherReportHandler(w http.ResponseWriter, r *http.Request)
}
