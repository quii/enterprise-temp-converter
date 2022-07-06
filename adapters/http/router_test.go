package http

import (
	"github.com/saltpay/enterprise-temp-converter/assert"
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"net/http"
	"testing"
)

func TestHTTPRouter(t *testing.T) {
	driver, stopServer := NewDriver()
	defer stopServer()

	t.Run("it passes the temp spec", func(t *testing.T) {
		specifications.ConvertTemperatures(t, driver)
	})

	t.Run("returns a bad request with a silly temp", func(t *testing.T) {
		res, err := http.Get(driver.URL + cToFPath + "?temp=lmao")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	})
}
