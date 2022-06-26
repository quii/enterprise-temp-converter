package main

import (
	"github.com/saltpay/enterprise-temp-converter/adapters/command_line"
	"testing"

	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestCommandLineTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	driver, err := command_line.NewCommandLineTempConverterDriver()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(driver.Cleanup)
	specifications.ItConvertsTemperatures(t, driver)
}
