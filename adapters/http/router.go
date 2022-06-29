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
	converter temperature.TempConverterSystem
}

func NewRouter(converter temperature.TempConverterSystem) http.Handler {
	svr := converterServer{converter: converter}

	mux := http.NewServeMux()
	mux.HandleFunc(cToFPath, svr.celsiusToFahrenHeit)
	mux.HandleFunc(fToCPath, svr.fahrenheitToCelsius)
	return mux
}

func (s converterServer) celsiusToFahrenHeit(w http.ResponseWriter, r *http.Request) {
	temp, err := getTempFromReq(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, _ := s.converter.ConvertFromCelsiusToFahrenheit(r.Context(), temp)

	fmt.Fprintf(w, "%.2f", f)
}

func (s converterServer) fahrenheitToCelsius(w http.ResponseWriter, r *http.Request) {
	temp, err := getTempFromReq(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c, _ := s.converter.ConvertFromFahrenheitToCelsius(r.Context(), temp)

	fmt.Fprintf(w, "%.2f", c)
}

func getTempFromReq(r *http.Request) (float64, error) {
	tempQS := r.URL.Query().Get("temp")
	temp, err := strconv.ParseFloat(tempQS, 64)
	return temp, err
}
