package main

import (
	"log"
	"os"

	"github.com/saltpay/enterprise-temp-converter/adapters/command_line"
	"github.com/saltpay/enterprise-temp-converter/cmd"
)

func main() {
	service, cleanUp, err := cmd.NewTemperatureConverterService()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	command_line.TempConverter(os.Stdin, os.Stdout, service)
}
