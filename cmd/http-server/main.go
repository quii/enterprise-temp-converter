package main

import (
	"log"
	"net/http"
	"os"

	temperature "github.com/saltpay/enterprise-temp-converter"
	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"github.com/saltpay/enterprise-temp-converter/telemetry"
)

func main() {
	converter := temperature.Converter{}
	loggingConverter := telemetry.NewLoggerMiddleware(os.Stderr, converter)

	if err := http.ListenAndServe(":8080", temphttp.NewRouter(loggingConverter)); err != nil {
		log.Fatal(err)
	}
}
