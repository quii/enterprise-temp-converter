package telemetry

import (
	"bytes"
	"context"
	"testing"

	temperature "github.com/saltpay/enterprise-temp-converter"
)

func TestNewLoggerMiddleware(t *testing.T) {
	var (
		out       bytes.Buffer
		converter temperature.Converter
		ctx       = context.Background()
	)

	logger := NewLoggerMiddleware(&out, converter)

	logger.ConvertFromCelsiusToFahrenheit(ctx, 20)
	wantedLog := "Converted 20.00 c to 68.00 f\n"
	if out.String() != wantedLog {
		t.Errorf("got %q, want %q", out.String(), wantedLog)
	}

	out.Reset()
	logger.ConvertFromFahrenheitToCelsius(ctx, 20)
	wantedLog = "Converted 20.00 f to -6.67 c\n"
	if out.String() != wantedLog {
		t.Errorf("got %q, want %q", out.String(), wantedLog)
	}
}
