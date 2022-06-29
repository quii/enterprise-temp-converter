package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestHTTPRouter(t *testing.T) {
	router := NewRouter(temperature.Converter{})
	server := httptest.NewServer(router)
	defer server.Close()

	driver := NewConverterHTTPClient(server.URL, &http.Client{Timeout: 2 * time.Second})
	specifications.ItConvertsTemperatures(t, driver)
}
