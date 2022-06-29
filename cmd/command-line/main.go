package main

import (
	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/adapters/command_line"
	"github.com/saltpay/enterprise-temp-converter/telemetry"

	"os"
)

func main() {
	converter := temperature.Converter{}
	loggingConverter := telemetry.NewLoggerMiddleware(os.Stderr, converter)

	command_line.TempConverter(os.Stdin, os.Stdout, loggingConverter)
}
