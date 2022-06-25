package main

import (
	temperature "github.com/saltpay/enterprise-temp-converter"
	"github.com/saltpay/enterprise-temp-converter/adapters/command_line"
	"os"
)

func main() {
	command_line.TempConverter(os.Stdin, os.Stdout, temperature.Converter{})
}
