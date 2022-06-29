package main

import (
	"testing"

	"github.com/saltpay/enterprise-temp-converter/cmd"
	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestCommandLineTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	cleanup, binPath, err := cmd.BuildBinary()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	driver := NewCommandLineTempConverterDriver(binPath)
	specifications.ItConvertsTemperatures(t, driver)
}
