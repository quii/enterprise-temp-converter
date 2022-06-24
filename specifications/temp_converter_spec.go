package specifications

import (
	"context"
	"testing"
)

type TempConverterSystem interface {
	ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error)
	ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error)
}

func ItConvertsTemperatures(t *testing.T, system TempConverterSystem) {
	t.Run("it converts from celsius to fahrenheit", func(t *testing.T) {
		var (
			celsius            = 32.0
			expectedFahrenheit = 89.6
			ctx                = context.Background()
		)

		actualFahrenheit, err := system.ConvertFromCelsiusToFahrenheit(ctx, celsius)

		if err != nil {
			t.Fatal(err)
		}

		if actualFahrenheit != expectedFahrenheit {
			t.Errorf("got %.2f, want %.2f", actualFahrenheit, expectedFahrenheit)
		}
	})

	t.Run("it converts from fahrenheit to celsius", func(t *testing.T) {
		var (
			expectedCelsius = 32.0
			fahrenheit      = 89.6
			ctx             = context.Background()
		)

		actualCelsius, err := system.ConvertFromFahrenheitToCelsius(ctx, fahrenheit)

		if err != nil {
			t.Fatal(err)
		}

		if actualCelsius != expectedCelsius {
			t.Errorf("got %.2f, want %.2f", actualCelsius, expectedCelsius)
		}
	})

}
