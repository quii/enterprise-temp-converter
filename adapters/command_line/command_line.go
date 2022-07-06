package command_line

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/saltpay/enterprise-temp-converter"
)

const promptText = `Press C to convert from Fahrenheit to Celsius.
Press F to convert from Celsius to Fahrenheit.
Your choice: `

func TempConverter(in io.Reader, out io.Writer, converterService temperature.TempConverterService) {
	fmt.Fprint(out, promptText)

	scanner := bufio.NewScanner(in)
	scanner.Scan()

	switch scanner.Text() {
	case "F":
		input, temp := promptForFloat(scanner, out, "Enter the temperature in Celsius")
		fahrenheit, _ := converterService.ConvertFromCelsiusToFahrenheit(context.Background(), temp)
		fmt.Fprintf(
			out,
			"The temperature %s in Fahrenheit is: %.2f\n",
			input,
			fahrenheit,
		)
	case "C":
		input, temp := promptForFloat(scanner, out, "Enter the temperature in Fahrenheit")
		celsius, _ := converterService.ConvertFromFahrenheitToCelsius(context.Background(), temp)
		fmt.Fprintf(
			out,
			"The temperature %s in Celsius is: %.2f\n",
			input,
			celsius,
		)
	}
}

func promptForFloat(scanner *bufio.Scanner, out io.Writer, prompt string) (string, float64) {
	fmt.Fprint(out, prompt+": ")
	scanner.Scan()
	f, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text(), f
}
