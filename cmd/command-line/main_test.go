package main

import (
	"testing"

	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestCommandLineTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	driver, err := NewCommandLineTempConverterDriver()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(driver.Cleanup)
	specifications.ItConvertsTemperatures(t, driver)
}
