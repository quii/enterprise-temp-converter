package main

import (
	"bytes"
	"context"
	"fmt"
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

	driver := http.NewConverterHTTPDriver(url, &http2.Client{})
	return driver.ConvertFromFahrenheitToCelsius(ctx, fahrenheit)
}

func (c *HTTPServerTempConverterDriver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	url, cleanup, err := c.runServer(ctx)
	if err != nil {
		return 0, err
	}
	defer cleanup()

	driver := http.NewConverterHTTPDriver(url, &http2.Client{})
	return driver.ConvertFromCelsiusToFahrenheit(ctx, celsius)
}

func (c *HTTPServerTempConverterDriver) runServer(ctx context.Context) (url string, cleanup func() error, err error) {
	cmd := exec.CommandContext(ctx, c.cmdPath)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return "", nil, fmt.Errorf("cannot run temp converter: %s", err)
	}
	time.Sleep(1 * time.Second) //todo: there will be a better way

	return "http://localhost:8080", cmd.Process.Kill, nil
}
