// Ex6: Retirement Calculator
//
// - Prompt the user to enter their current age and desired retirement age.
// - Convert input to numeric values before performing calculations.
// - Determine how many years are left until retirement.
// - Get the current year from the system, not hard-coded.
// - Calculate and display the retirement year and remaining years.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Prompt the user to enter their current age and desired retirement age
	var currentAge, retirementAge int
	fmt.Print("What is your current age? ")
	fmt.Scanln(&currentAge)
	fmt.Print("At what age would you like to retire? ")
	fmt.Scanln(&retirementAge)

	// Calculate how many years are left until retirement
	yearsLeft := retirementAge - currentAge

	// Get the current year from the system
	currentYear := time.Now().Year()

	// Calculate the retirement year
	retirementYear := currentYear + yearsLeft

	// Display the results
	fmt.Printf("You have %d years left until you can retire.\n", yearsLeft)
	fmt.Printf("It's %d, so you can retire in %d.\n", currentYear, retirementYear)
}
