package specifications

import (
	"context"
	"testing"

	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/assert"
)

func ItConvertsTemperatures(t *testing.T, driver temperature.TempConverterService) {
	t.Run("it converts from celsius to fahrenheit", func(t *testing.T) {
		var (
			celsius            = 32.0
			expectedFahrenheit = 89.6
			ctx                = context.Background()
		)

		actualFahrenheit, err := driver.ConvertFromCelsiusToFahrenheit(ctx, celsius)
		assert.NoError(t, err)
		assert.Equal(t, actualFahrenheit, expectedFahrenheit)
	})

	t.Run("it converts from fahrenheit to celsius", func(t *testing.T) {
		var (
			expectedCelsius = 32.0
			fahrenheit      = 89.6
			ctx             = context.Background()
		)

		actualCelsius, err := driver.ConvertFromFahrenheitToCelsius(ctx, fahrenheit)
		assert.NoError(t, err)
		assert.Equal(t, actualCelsius, expectedCelsius)
	})

}
