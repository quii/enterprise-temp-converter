package temperature

func ConvertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func ConvertCelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
