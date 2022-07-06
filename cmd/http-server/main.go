package main

import (
	"log"
	"net/http"

	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"github.com/saltpay/enterprise-temp-converter/cmd"
)

const port = ":8080"

func main() {
	service, cleanUp, err := cmd.NewTemperatureConverterService()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	router := temphttp.NewRouter(service)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
