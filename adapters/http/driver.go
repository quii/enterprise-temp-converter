package http

import (
	temperature "github.com/saltpay/enterprise-temp-converter"
	"net/http"
	"net/http/httptest"
	"time"
)

type Driver struct {
	*ConverterHTTPClient
	URL string
}

func NewDriver() (*Driver, func()) {
	router := NewRouter(temperature.Service{})
	server := httptest.NewServer(router)
	client := NewConverterHTTPClient(server.URL, &http.Client{Timeout: 2 * time.Second})
	return &Driver{
		client,
		server.URL,
	}, server.Close
}
