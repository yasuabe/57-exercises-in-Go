// Ex14: Tax Calculator
//
// - Prompt the user for the order amount and the state.
// - If the state is "WI", calculate 5.5% tax and display subtotal, tax, and total.
// - For other states, display only the total.
// - Use only a simple if statement (no else clause).
// - Round all money up to the nearest cent.
// - Use a single output statement at the end.
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

	// Prompt for order amount
	fmt.Print("What is the order amount? ")
	orderAmountInput, _ := reader.ReadString('\n')
	orderAmountInput = strings.TrimSpace(orderAmountInput)
	orderAmount, _ := strconv.ParseFloat(orderAmountInput, 64)

	// Prompt for state
	fmt.Print("What is the state? ")
	state, _ := reader.ReadString('\n')
	state = strings.TrimSpace(state)

	var taxRate, tax, total float64
	var output string

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
