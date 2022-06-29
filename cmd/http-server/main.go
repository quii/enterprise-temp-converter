package main

import (
	"log"
	"net/http"

	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
)

func main() {
	if err := http.ListenAndServe(":8080", temphttp.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
