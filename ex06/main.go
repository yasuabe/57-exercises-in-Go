// Ex6: Retirement Calculator
//
// - Prompt the user to enter their current age and desired retirement age.
// - Convert input to numeric values before performing calculations.
// - Determine how many years are left until retirement.
// - Get the current year from the system, not hard-coded.
// - Calculate and display the retirement year and remaining years.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter their current age
	fmt.Print("What is your current age? ")
	currentAgeInput, _ := reader.ReadString('\n')
	currentAgeInput = strings.TrimSpace(currentAgeInput)
	currentAge, _ := strconv.Atoi(currentAgeInput)

	// Prompt the user to enter their desired retirement age
	fmt.Print("At what age would you like to retire? ")
	retirementAgeInput, _ := reader.ReadString('\n')
	retirementAgeInput = strings.TrimSpace(retirementAgeInput)
	retirementAge, _ := strconv.Atoi(retirementAgeInput)

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
