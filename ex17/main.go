// # Ex17: Blood Alcohol Calculator
//
// - Prompt for weight, gender, alcohol amount, and time since last drink.
// - Compute BAC using a given formula.
// - Report whether it's legal to drive (BAC ≥ 0.08 means illegal).
// - Constraint: Validate that inputs are numeric.

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

	// Prompt for weight
	fmt.Print("Enter your weight (in pounds): ")
	weightInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	weightInput = strings.TrimSpace(weightInput)
	weight, err := strconv.ParseFloat(weightInput, 64)
	if err != nil || weight <= 0 {
		fmt.Println("Invalid input. Please enter a valid numeric weight.")
		return
	}

	// Prompt for gender
	fmt.Print("Enter your gender (male/female): ")
	genderInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	gender := strings.TrimSpace(strings.ToLower(genderInput))

	// Determine alcohol distribution ratio
	var r float64
	if gender == "male" {
		r = 0.73
	} else if gender == "female" {
		r = 0.66
	} else {
		fmt.Println("Invalid input. Please enter 'male' or 'female' for gender.")
		return
	}

	// Prompt for total alcohol consumed
	fmt.Print("Enter the total alcohol consumed (in ounces): ")
	alcoholInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	alcoholInput = strings.TrimSpace(alcoholInput)
	alcohol, err := strconv.ParseFloat(alcoholInput, 64)
	if err != nil || alcohol < 0 {
		fmt.Println("Invalid input. Please enter a valid numeric value for alcohol consumed.")
		return
	}

	// Prompt for hours since last drink
	fmt.Print("Enter the number of hours since your last drink: ")
	hoursInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	hoursInput = strings.TrimSpace(hoursInput)
	hours, err := strconv.ParseFloat(hoursInput, 64)
	if err != nil || hours < 0 {
		fmt.Println("Invalid input. Please enter a valid numeric value for hours.")
		return
	}

	// Calculate BAC
	bac := (alcohol * 5.14 / weight * r) - (0.15 * hours)

	// Output BAC and driving legality
	fmt.Printf("Your BAC is %.2f\n", bac)
	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}
