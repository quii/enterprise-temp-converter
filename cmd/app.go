package cmd

import (
	"os"
	"time"

	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/telemetry"
)

func NewApp() (temperature.TempConverterSystem, func() error, error) {
	f, err := os.OpenFile("../log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, nil, err
	}

	converter := temperature.Converter{}
	loggingConverter := telemetry.NewLoggerMiddleware(f, converter, time.Now)

	return loggingConverter, f.Close, nil
}
