package main

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
	binName := baseBinName

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		return nil, fmt.Errorf("cannot build tool %s: %s", binName, err)
	}

	build.Wait()

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cmdPath := filepath.Join(dir, binName)

	return &CommandLineTempConverterDriver{
		binName: binName,
		cmdPath: cmdPath,
	}, nil
}

func (c *CommandLineTempConverterDriver) ConvertFromCelsiusToFahrenheit(ctx context.Context, celsius float64) (fahrenheit float64, err error) {
	cmd := exec.Command(c.cmdPath)
	cmdStdin, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("cannot run temp converter: %s", err)
	}
	fmt.Fprintln(cmdStdin, "F")
	fmt.Fprintln(cmdStdin, fmt.Sprintf("%.2f", celsius))
	cmd.Wait()

	lines := strings.Split(out.String(), "\n")
	lastLine := lines[len(lines)-2]
	indexFunc := strings.LastIndexFunc(lastLine, func(r rune) bool {
		return r == ':'
	})
	f := lastLine[indexFunc+2:]
	fahrenheit, err = strconv.ParseFloat(f, 2)
	if err != nil {
		return 0.0, err
	}

	return fahrenheit, err
}

func (c *CommandLineTempConverterDriver) Cleanup() {
	os.Remove(c.binName)
}
