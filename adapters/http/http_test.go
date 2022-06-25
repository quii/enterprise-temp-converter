package http

import (
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPServer(t *testing.T) {
	router := NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	driver := NewConverterDriver(server.URL, &http.Client{Timeout: 2 * time.Second})
	specifications.ItConvertsTemperatures(t, driver)
}
