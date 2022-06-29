package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ConverterDriver struct {
	baseURL string
	client  *http.Client
}

func NewConverterHTTPDriver(baseURL string, client *http.Client) *ConverterDriver {
	return &ConverterDriver{baseURL: baseURL, client: client}
}

func (c *ConverterDriver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	return c.convertATemp(ctx, celsius, cToFPath)
}

func (c *ConverterDriver) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	return c.convertATemp(ctx, fahrenheit, fToCPath)
}

func (c *ConverterDriver) convertATemp(ctx context.Context, in float64, path string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	q := req.URL.Query()
	q.Add("temp", fmt.Sprintf("%.2f", in))
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return 0, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	s := string(body)
	return strconv.ParseFloat(s, 64)
}
