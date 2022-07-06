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
		in  = strings.NewReader(fmt.Sprintf("F\n%.2f", celsius))
		out bytes.Buffer
	)

	TempConverter(in, &out, temperature.Service{})
	return extractTempFromStdout(out)
}

func (d Driver) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	var (
		in  = strings.NewReader(fmt.Sprintf("C\n%.2f", fahrenheit))
		out bytes.Buffer
	)

	TempConverter(in, &out, temperature.Service{})
	return extractTempFromStdout(out)
}

func extractTempFromStdout(out bytes.Buffer) (float64, error) {
	output := out.String()
	resultTxt := strings.TrimSpace(output[len(output)-7:])
	return strconv.ParseFloat(resultTxt, 64)
}
