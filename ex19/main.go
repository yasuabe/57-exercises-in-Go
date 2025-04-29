// # Ex19:	BMI Calculator
//
// - Prompt the user for their height (in inches) and weight (in pounds).
// - Calculate BMI using:
// - ` bmi = (weight / (height × height)) × 703 ``
// - Display:
//   - “You are within the ideal weight range.” if BMI is 18.5–25.
//   - Otherwise, indicate if the user is underweight or overweight and advise seeing a doctor.
// - Input must be numeric—reject non-numeric input.

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

	height := getValidInput(reader, "Enter your height in inches: ")
	weight := getValidInput(reader, "Enter your weight in pounds: ")

	bmi := (weight / (height * height)) * 703

	fmt.Printf("Your BMI is %.1f.\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are underweight. You should see your doctor.")
	} else if bmi <= 25 {
		fmt.Println("You are within the ideal weight range.")
	} else {
		fmt.Println("You are overweight. You should see your doctor.")
	}
}

func getValidInput(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value > 0 {
			return value
		}

		fmt.Println("Invalid input. Please enter a valid numeric value.")
	}
}
