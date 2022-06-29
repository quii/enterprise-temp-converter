package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	baseBinName = "temp-converter-testbinary"
)

func BuildBinary() (cleanup func(), cmdPath string, err error) {
	binName := baseBinName

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		return nil, "", fmt.Errorf("cannot build tool %s: %s", binName, err)
	}

	build.Wait()

	dir, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	cmdPath = filepath.Join(dir, binName)

	cleanup = func() {
		os.Remove(binName)
	}

	return
}
