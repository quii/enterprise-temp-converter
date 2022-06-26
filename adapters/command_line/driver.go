package command_line

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	baseBinName = "temp-converter-testbinary"
)

type CommandLineTempConverterDriver struct {
	binName string
	cmdPath string
}

func NewCommandLineTempConverterDriver() (*CommandLineTempConverterDriver, error) {
	binName, cmdPath, err := buildBinary()

	if err != nil {
		return nil, err
	}

	return &CommandLineTempConverterDriver{
		binName: binName,
		cmdPath: cmdPath,
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
	os.Remove(c.binName)
}

func buildBinary() (binName, cmdPath string, err error) {
	binName = baseBinName

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		return "", "", fmt.Errorf("cannot build tool %s: %s", binName, err)
	}

	build.Wait()

	dir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	cmdPath = filepath.Join(dir, binName)

	return
}
