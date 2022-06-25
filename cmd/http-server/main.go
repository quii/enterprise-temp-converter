package main

import (
	temphttp "github.com/saltpay/enterprise-temp-converter/adapters/http"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", temphttp.NewRouter())
}
