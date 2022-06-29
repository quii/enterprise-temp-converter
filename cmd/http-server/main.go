package main

import (
	"log"
	"net/http"

	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"github.com/saltpay/enterprise-temp-converter/cmd"
)

func main() {
	converter, cleanUp, err := cmd.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	router := temphttp.NewRouter(converter)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
