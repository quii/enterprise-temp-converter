package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	var temp string
	fmt.Scanln(&temp)
	fmt.Printf("You chose %q\n", temp)
	switch temp {
	case "F":
		fmt.Print("Enter the temperature in Celsius: ")
		var in string
		fmt.Scanln(&in)
		c, err := strconv.ParseFloat(in, 2)
		if err != nil {
			log.Fatal(err)
		}
		fahrenheit := (c * 9 / 5) + 32
		fmt.Printf("The temperature %s in Fahrenheit is: %.2f\n", in, fahrenheit)
	case "C":
		fmt.Print("Enter the temperature in Fahrenheit: ")
		var in string
		fmt.Scanln(&in)
		f, err := strconv.ParseFloat(in, 2)
		if err != nil {
			log.Fatal(err)
		}
		celsius := (f - 32) * 5 / 9
		fmt.Printf("The temperature %s in Celsius is: %.2f\n", in, celsius)
	}
}
