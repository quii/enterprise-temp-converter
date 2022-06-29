package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/assert"
	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestHTTPRouter(t *testing.T) {
	router := NewRouter(temperature.Converter{})
	server := httptest.NewServer(router)
	defer server.Close()

	t.Run("it passes the temp spec", func(t *testing.T) {
		driver := NewConverterHTTPClient(server.URL, &http.Client{Timeout: 2 * time.Second})
		specifications.ItConvertsTemperatures(t, driver)
	})

	t.Run("returns a bad request with a silly temp", func(t *testing.T) {
		res, err := http.Get(server.URL + cToFPath + "?temp=lmao")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	})
}
