package temperature

import (
	"context"
)

type TempConverterService interface {
	ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error)
	ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error)
}

type Service struct {
}

func (c Service) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	return (celsius * 9 / 5) + 32, nil
}

func (c Service) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	return (fahrenheit - 32) * 5 / 9, nil
}
