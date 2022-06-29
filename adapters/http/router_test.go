package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestHTTPRouter(t *testing.T) {
	router := NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	driver := NewConverterHTTPDriver(server.URL, &http.Client{Timeout: 2 * time.Second})
	specifications.ItConvertsTemperatures(t, driver)
}
