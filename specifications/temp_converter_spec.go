package specifications

import (
	"context"
	"testing"
)

type TempConverterSystem interface {
	ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error)
}

func ItConvertsTempsFromCtoF(t *testing.T, system TempConverterSystem) {
	t.Run("it converts from celsius to fahrenheit", func(t *testing.T) {
		var (
			celsius            = 0.0
			expectedFahrenheit = 32.0
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
}
