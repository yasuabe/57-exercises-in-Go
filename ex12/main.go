// Ex12: Computing Simple Interest
//
// - Prompt for principal, interest rate (as %), and years.
// - Compute simple interest: A = P × (1 + r × t).
// - Convert percent rate by dividing by 100.
// - Round up to the nearest cent.
// - Format the output as currency.

package main

import "fmt"

func main() {
	// Prompt for principal, interest rate, and years
	var principal, rate, years float64

	// Get user input
	fmt.Print("Enter the principal: ")
	fmt.Scanln(&principal)
	fmt.Print("Enter the rate of interest: ")
	fmt.Scanln(&rate)
	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)

	// Compute simple interest
	rate = rate / 100 // Convert percent rate to decimal
	amount := principal * (1 + rate*years)

	// Round up to the nearest cent
	amount = float64(int(amount*100+0.999999)) / 100.0

	// Print result formatted as currency
	fmt.Printf("After %.0f years at %.1f%%, the investment will\nbe worth $%.2f.\n", years, rate*100, amount)
}
