package main

import (
	http2 "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPServer(t *testing.T) {
	router := http2.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	driver := http2.NewConverterDriver(server.URL, &http.Client{Timeout: 2 * time.Second})
	specifications.ItConvertsTemperatures(t, driver)
}
