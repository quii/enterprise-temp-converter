package main

import (
	"context"
	"fmt"
	"log"
	"net"
	http2 "net/http"
	"os/exec"
	"time"

	"github.com/saltpay/enterprise-temp-converter/adapters/http"
)

type HTTPServerTempConverterDriver struct {
	cmdPath string
}

func NewHTTPServerTempConverterDriver(cmdPath string) *HTTPServerTempConverterDriver {
	return &HTTPServerTempConverterDriver{
		cmdPath: cmdPath,
	}
}

func (c *HTTPServerTempConverterDriver) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	url, cleanup, err := c.runServer(ctx)
	if err != nil {
		return 0, err
	}
	defer cleanup()

	client := http.NewConverterHTTPClient(url, &http2.Client{})
	return client.ConvertFromFahrenheitToCelsius(ctx, fahrenheit)
}

func (c *HTTPServerTempConverterDriver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	url, cleanup, err := c.runServer(ctx)
	if err != nil {
		return 0, err
	}
	defer cleanup()

	driver := http.NewConverterHTTPClient(url, &http2.Client{})
	return driver.ConvertFromCelsiusToFahrenheit(ctx, celsius)
}

func (c *HTTPServerTempConverterDriver) runServer(ctx context.Context) (url string, cleanup func() error, err error) {
	cmd := exec.CommandContext(ctx, c.cmdPath)

	if err := cmd.Start(); err != nil {
		return "", nil, fmt.Errorf("cannot run temp converter: %s", err)
	}
	waitForServerListening()

	return "http://localhost:8080", cmd.Process.Kill, nil
}

func waitForServerListening() {
	for i := 0; i < 20; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", "8080"))
		if conn != nil {
			log.Println("a connection!")
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
