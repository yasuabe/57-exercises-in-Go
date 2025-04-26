// Ex13: Determining Compound Interest
//
// - Prompt the user for principal amount, interest rate (as a percentage), number of years, and compounding frequency per year.
// - Convert the interest rate by dividing it by 100.
// - Use the compound interest formula to compute the final amount.
// - Round up fractions of a cent to the next penny.
// - Format the output as money.
package main

import (
	"fmt"
	"math"
)

func main() {
	// Prompt for principal, interest rate, years, and compounding frequency
	var principal, rate float64
	var times, years int

	// Get user input
	fmt.Print("What is the principal amount? ")
	fmt.Scanln(&principal)
	fmt.Print("What is the rate? ")
	fmt.Scanln(&rate)
	fmt.Print("What is the number of years? ")
	fmt.Scanln(&years)
	fmt.Print("What is the number of times the interest is compounded per year? ")
	fmt.Scanln(&times)

	// Convert percent rate to decimal
	rate = rate / 100

	// Compute compound interest using the correct formula
	amount := principal * math.Pow(1+rate/float64(times), float64(times*years))

	// Round up to the nearest cent
	amount = float64(int(amount*100+0.999999)) / 100.0 // TODO: duplication

	// Print result formatted as currency
	fmt.Printf(
		"$%.2f invested at %.1f%% for %d years compounded %d times per year is $%.2f.\n",
		principal,
		rate*100,
		years,
		times,
		amount)
}
