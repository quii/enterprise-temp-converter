package command_line

import (
	"context"
	"fmt"
	"github.com/saltpay/enterprise-temp-converter"
	"io"
	"log"
	"strconv"
)

func TempConverter(in io.Reader, out io.Writer, converter temperature.Converter) {
	fmt.Fprintln(out, "Press C to convert from Fahrenheit to Celsius.")
	fmt.Fprintln(out, "Press F to convert from Celsius to Fahrenheit.")
	fmt.Fprintln(out, "Your choice: ")
	var temp string
	fmt.Fscanln(in, &temp)

	switch temp {
	case "F":
		input, temp := promptForFloat(in, "Enter the temperature in Celsius")
		fahrenheit, _ := converter.ConvertFromCelsiusToFahrenheit(context.Background(), temp)
		fmt.Fprintf(
			out,
			"The temperature %s in Fahrenheit is: %.2f\n",
			input,
			fahrenheit,
		)
	case "C":
		input, temp := promptForFloat(in, "Enter the temperature in Fahrenheit")
		celsius, _ := converter.ConvertFromFahrenheitToCelsius(context.Background(), temp)
		fmt.Fprintf(
			out,
			"The temperature %s in Celsius is: %.2f\n",
			input,
			celsius,
		)
	}
}

func promptForFloat(in io.Reader, prompt string) (string, float64) {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Fscanln(in, &input)
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return input, f
}