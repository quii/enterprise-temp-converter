package temperature_test

import (
	"testing"

	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/specifications"
)

func TestCommandLineTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	specifications.ConvertTemperatures(t, temperature.Service{})
}
