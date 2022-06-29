package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/saltpay/enterprise-temp-converter/cmd"
)

type CommandLineTempConverterDriver struct {
	cmdPath string
	cleanup func()
}

func NewCommandLineTempConverterDriver() (*CommandLineTempConverterDriver, error) {
	cleanup, cmdPath, err := cmd.BuildBinary()

	if err != nil {
		return nil, err
	}

	return &CommandLineTempConverterDriver{
		cmdPath: cmdPath,
		cleanup: cleanup,
	}, nil
}

func (c *CommandLineTempConverterDriver) ConvertFromFahrenheitToCelsius(ctx context.Context, fahrenheit float64) (celsius float64, err error) {
	return c.runProgram(ctx, "C", fahrenheit)
}

func (c *CommandLineTempConverterDriver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	return c.runProgram(ctx, "F", celsius)
}

func (c *CommandLineTempConverterDriver) runProgram(ctx context.Context, choice string, temp float64) (conversion float64, err error) {
	cmd := exec.CommandContext(ctx, c.cmdPath)
	cmdStdin, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("cannot run temp converter: %s", err)
	}
	fmt.Fprintln(cmdStdin, choice)
	fmt.Fprintln(cmdStdin, fmt.Sprintf("%.2f", temp))
	cmd.Wait()

	lines := strings.Split(out.String(), "\n")
	lastLine := lines[len(lines)-2]
	indexFunc := strings.LastIndexFunc(lastLine, func(r rune) bool {
		return r == ':'
	})
	result, err := strconv.ParseFloat(lastLine[indexFunc+2:], 64)
	if err != nil {
		return 0.0, err
	}

	return result, err
}

func (c *CommandLineTempConverterDriver) Cleanup() {
	c.cleanup()
}
