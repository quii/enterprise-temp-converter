package telemetry

import (
	"bytes"
	"context"
	"testing"
	"time"

	temperature "github.com/saltpay/enterprise-temp-converter"
)

func TestNewLoggerMiddleware(t *testing.T) {
	var (
		out bytes.Buffer
		now = func() time.Time {
			return time.Date(1984, 7, 5, 8, 0, 0, 0, time.UTC)
		}
		converter temperature.Converter
		ctx       = context.Background()
	)

	logger := NewLoggerMiddleware(&out, converter, now)

	logger.ConvertFromCelsiusToFahrenheit(ctx, 20)
	wantedLog := "1984-07-05T08:00:00Z Converted 20.00 c to 68.00 f\n"
	if out.String() != wantedLog {
		t.Errorf("got %q, want %q", out.String(), wantedLog)
	}

	out.Reset()
	logger.ConvertFromFahrenheitToCelsius(ctx, 20)
	wantedLog = "1984-07-05T08:00:00Z Converted 20.00 f to -6.67 c\n"
	if out.String() != wantedLog {
		t.Errorf("got %q, want %q", out.String(), wantedLog)
	}
}
