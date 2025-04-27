// Ex14: Tax Calculator
//
// - Prompt the user for the order amount and the state.
// - If the state is "WI", calculate 5.5% tax and display subtotal, tax, and total.
// - For other states, display only the total.
// - Use only a simple if statement (no else clause).
// - Round all money up to the nearest cent.
// - Use a single output statement at the end.
package main

import "fmt"

func main() {
	// Prompt for order amount and state
	var orderAmount, taxRate, tax, total float64
	var state, output string

	// Get user input
	fmt.Print("What is the order amount? ")
	fmt.Scanln(&orderAmount)
	fmt.Print("What is the state? ")
	fmt.Scanln(&state)

	// Check if the state is WI
	if state == "WI" {
		// Set tax rate for Wisconsin
		taxRate = 0.055
		// Calculate tax
		tax = orderAmount * taxRate
		// Round tax to the nearest cent
		tax = float64(int(tax*100+0.999999)) / 100.0
		// Prepare output for WI
		output = fmt.Sprintf("The subtotal is $%.2f.\nThe tax is $%.2f.\n", orderAmount, tax)
	}

	// Calculate total (always includes tax, which is 0 if state is not WI)
	total = orderAmount + tax
	// Round total to the nearest cent
	total = float64(int(total*100+0.999999)) / 100.0

	// Append the total to the output
	output += fmt.Sprintf("The total is $%.2f.\n", total)

	// Use a single output statement to display the results
	fmt.Print(output)
}
