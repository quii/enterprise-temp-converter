package main

import "fmt"

func main() {
	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	var temp string
	fmt.Scanln(&temp)
	fmt.Printf("You chose %q\n", temp)
	switch temp {
	case "F":
		fmt.Println("Enter the temperature in Fahrenheit: ")
		var f string
		fmt.Scanln(&f)
		fmt.Printf("The temperature %s in Celsius is: 32\n", f)
	}
}
