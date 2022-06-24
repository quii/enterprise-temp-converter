package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/saltpay/enterprise-temp-converter"
)

func main() {
	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	var temp string
	fmt.Scanln(&temp)

	switch temp {
	case "F":
		in, temp := PromptForFloat("Enter the temperature in Celsius")
		fmt.Printf(
			"The temperature %s in Fahrenheit is: %.2f\n",
			in,
			temperature.ConvertCelsiusToFahrenheit(temp),
		)
	case "C":
		in, temp := PromptForFloat("Enter the temperature in Fahrenheit")
		fmt.Printf(
			"The temperature %s in Celsius is: %.2f\n",
			in,
			temperature.ConvertFahrenheitToCelsius(temp),
		)
	}
}

func PromptForFloat(prompt string) (string, float64) {
	fmt.Print(prompt + ": ")
	var in string
	fmt.Scanln(&in)
	f, err := strconv.ParseFloat(in, 2)
	if err != nil {
		log.Fatal(err)
	}
	return in, f
}
