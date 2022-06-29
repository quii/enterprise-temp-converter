package http

import (
	"fmt"
	"net/http"
	"strconv"

	temperature "github.com/saltpay/enterprise-temp-converter"
)

const (
	cToFPath = "/celsius-to-fahrenheit"
	fToCPath = "/fahrenheit-to-celsius"
)

type converterServer struct {
	converter temperature.Converter
}

func NewRouter() http.Handler {
	svr := converterServer{converter: temperature.Converter{}}

	mux := http.NewServeMux()
	mux.HandleFunc(cToFPath, svr.celsiusToFahrenHeit)
	mux.HandleFunc(fToCPath, svr.fahrenheitToCelsius)
	return mux
}

func (s converterServer) celsiusToFahrenHeit(w http.ResponseWriter, r *http.Request) {
	tempQS := r.URL.Query().Get("temp")
	temp, _ := strconv.ParseFloat(tempQS, 64)
	f, _ := s.converter.ConvertFromCelsiusToFahrenheit(r.Context(), temp)

	fmt.Fprintf(w, "%.2f", f)
}

func (s converterServer) fahrenheitToCelsius(w http.ResponseWriter, r *http.Request) {
	tempQS := r.URL.Query().Get("temp")
	temp, _ := strconv.ParseFloat(tempQS, 64)
	c, _ := s.converter.ConvertFromFahrenheitToCelsius(r.Context(), temp)

	fmt.Fprintf(w, "%.2f", c)
}
