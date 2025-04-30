// # Ex20:
//
// - Prompt for order amount and shipping state.
// - For Wisconsin: also ask for county and apply 5% + extra tax based on county.
//   - For Eau Claire county residents, add an additional 0.005 tax.
//   - For Dunn county residents, add an additional 0.004 tax.
// - For Illinois: apply a flat 8% tax.
// - For other states: no tax.
// - Round all money up to the nearest cent.
// - Constraint: Output both tax and total (if taxed) using a single print statement.

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

	// Prompt for the order amount
	fmt.Print("What is the order amount? ")
	orderInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	orderInput = strings.TrimSpace(orderInput)
	orderAmount, err := strconv.ParseFloat(orderInput, 64)
	if err != nil || orderAmount < 0 {
		fmt.Println("Invalid order amount. Please enter a valid number.")
		return
	}

	// Prompt for the state
	fmt.Print("What state do you live in? ")
	stateInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	state := strings.TrimSpace(strings.ToLower(stateInput))

	var taxRate float64
	var county string

	// Handle state-specific tax rules
	if state == "wisconsin" {
		taxRate = 0.05 // Base tax for Wisconsin

		// Prompt for the county
		fmt.Print("Enter the county: ")
		countyInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			return
		}
		county = strings.TrimSpace(strings.ToLower(countyInput))

		// Add county-specific tax
		if county == "eau claire" {
			taxRate += 0.005
		} else if county == "dunn" {
			taxRate += 0.004
		}
	} else if state == "illinois" {
		taxRate = 0.08 // Flat tax for Illinois
	}

	// Calculate tax and total using integer arithmetic to avoid floating-point errors
	tax := int(orderAmount*taxRate*100 + 0.9999) // Always round up to the nearest cent
	total := int(orderAmount*100) + tax          // Add tax to the order amount in cents

	// Convert back to dollars for output
	taxDollars := float64(tax) / 100
	totalDollars := float64(total) / 100

	// Output the results
	if taxRate > 0 {
		fmt.Printf("The tax is $%.2f.\nThe total is $%.2f.\n", taxDollars, totalDollars)
	} else {
		fmt.Printf("The total is $%.2f.\n", orderAmount)
	}
}
