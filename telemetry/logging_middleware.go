package telemetry

import (
	"context"
	"fmt"
	"io"
	"time"

	temperature "github.com/saltpay/enterprise-temp-converter"
)

type LoggerMiddleware struct {
	out      io.Writer
	delegate temperature.TempConverterService
	now      func() time.Time
}

func NewLoggerMiddleware(out io.Writer, delegate temperature.TempConverterService, now func() time.Time) *LoggerMiddleware {
	return &LoggerMiddleware{out: out, delegate: delegate, now: now}
}

func (l LoggerMiddleware) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	f, err := l.delegate.ConvertFromCelsiusToFahrenheit(ctx, celsius)
	if err != nil {
		return 0, err
	}
	fmt.Fprintf(l.out, "%s Converted %.2f c to %.2f f\n", l.dateAndTime(), celsius, f)
	return f, nil
}

func (l LoggerMiddleware) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	c, err := l.delegate.ConvertFromFahrenheitToCelsius(ctx, fahrenheit)
	if err != nil {
		return 0, err
	}
	fmt.Fprintf(l.out, "%s Converted %.2f f to %.2f c\n", l.dateAndTime(), fahrenheit, c)
	return c, nil
}

func (l LoggerMiddleware) dateAndTime() string {
	return l.now().UTC().Format(time.RFC3339)
}
