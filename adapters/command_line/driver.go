package command_line

import (
	"bytes"
	"context"
	"fmt"
	"github.com/saltpay/enterprise-temp-converter"
	"strconv"
	"strings"
)

type Driver struct {
}

func (d Driver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	var (
		in  bytes.Buffer
		out bytes.Buffer
	)

	fmt.Fprintln(&in, convertToFahrenheitChoice)
	fmt.Fprintln(&in, celsius)

	TempConverter(&in, &out, temperature.Service{})
	return extractTempFromStdout(&out)
}

func (d Driver) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	var (
		in  bytes.Buffer
		out bytes.Buffer
	)

	fmt.Fprintln(&in, convertToCelsiusChoice)
	fmt.Fprintln(&in, fahrenheit)

	TempConverter(&in, &out, temperature.Service{})
	return extractTempFromStdout(&out)
}

func extractTempFromStdout(out *bytes.Buffer) (float64, error) {
	output := out.String()
	resultTxt := strings.TrimSpace(output[len(output)-7:])
	return strconv.ParseFloat(resultTxt, 64)
}
