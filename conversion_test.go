package temperature_test

import (
	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"testing"
)

func TestCommandLineTempConverterDriver_ConvertFromCelsiusToFahrenheit(t *testing.T) {
	converter := temperature.Converter{}
	specifications.ItConvertsTemperatures(t, converter)
}
