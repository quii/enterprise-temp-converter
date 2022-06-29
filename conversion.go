package temperature

import (
	"context"
)

type TempConverterSystem interface {
	ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error)
	ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error)
}

type Converter struct {
}

func (c Converter) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	return (celsius * 9 / 5) + 32, nil
}

func (c Converter) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	return (fahrenheit - 32) * 5 / 9, nil
}
