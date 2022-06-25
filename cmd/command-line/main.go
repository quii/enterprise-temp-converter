package main

import (
	"context"
	"fmt"
	temperature "github.com/saltpay/enterprise-temp-converter"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	var temp string
	fmt.Scanln(&temp)

	converter := temperature.Converter{}

	switch temp {
	case "F":
		in, temp := PromptForFloat("Enter the temperature in Celsius")
		fahrenheit, _ := converter.ConvertFromCelsiusToFahrenheit(context.Background(), temp)
		fmt.Printf(
			"The temperature %s in Fahrenheit is: %.2f\n",
			in,
			fahrenheit,
		)
	case "C":
		in, temp := PromptForFloat("Enter the temperature in Fahrenheit")
		celsius, _ := converter.ConvertFromFahrenheitToCelsius(context.Background(), temp)
		fmt.Printf(
			"The temperature %s in Celsius is: %.2f\n",
			in,
			celsius,
		)
	}
}

func PromptForFloat(prompt string) (string, float64) {
	fmt.Print(prompt + ": ")
	var in string
	fmt.Scanln(&in)
	f, err := strconv.ParseFloat(in, 64)
	if err != nil {
		log.Fatal(err)
	}
	return in, f
}
