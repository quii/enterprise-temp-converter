package main

import (
	"log"
	"net/http"

	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"github.com/saltpay/enterprise-temp-converter/cmd"
)

func main() {
	service, cleanUp, err := cmd.NewTemperatureConverterService()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	router := temphttp.NewRouter(service)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
