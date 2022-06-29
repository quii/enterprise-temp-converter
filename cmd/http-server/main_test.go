package main

import (
	"testing"

	"github.com/saltpay/enterprise-temp-converter/cmd"
	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestHTTPServerTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	cleanup, binPath, err := cmd.BuildBinary()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	driver := NewHTTPServerTempConverterDriver(binPath)
	specifications.ItConvertsTemperatures(t, driver)
}
