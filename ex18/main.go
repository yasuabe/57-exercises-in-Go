// # Ex18:	Temperature Converter
//
// - Prompt the user to choose conversion type: Fahrenheit ↔ Celsius.
// - Accept both uppercase and lowercase (C, F).
// - Prompt for the input temperature based on the choice.
// - Convert using the appropriate formula.
// - Display the result using minimal and non-redundant output statements.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for conversion type
	var choice string
	for {
		fmt.Println("Press C to convert from Fahrenheit to Celsius.")
		fmt.Println("Press F to convert from Celsius to Fahrenheit.")
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToUpper(input))
		if choice == "C" || choice == "F" {
			break
		}
		fmt.Println("Invalid choice. Please enter C or F.")
	}

	// Prompt for temperature
	fmt.Print("Please enter the temperature: ")
	tempInput, _ := reader.ReadString('\n')
	temp, err := strconv.ParseFloat(strings.TrimSpace(tempInput), 64)
	if err != nil {
		fmt.Println("Invalid temperature input.")
		return
	}

	// Perform conversion
	var result float64
	if choice == "C" {
		result = (temp - 32) * 5 / 9
		fmt.Printf("The temperature in Celsius is %.1f.\n", result)
	} else {
		result = (temp * 9 / 5) + 32
		fmt.Printf("The temperature in Fahrenheit is %.1f.\n", result)
	}
}
