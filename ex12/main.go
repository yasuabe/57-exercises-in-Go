// Ex12: Computing Simple Interest
//
// - Prompt for principal, interest rate (as %), and years.
// - Compute simple interest: A = P × (1 + r × t).
// - Convert percent rate by dividing by 100.
// - Round up to the nearest cent.
// - Format the output as currency.

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

	// Prompt for principal
	fmt.Print("Enter the principal: ")
	principalInput, _ := reader.ReadString('\n')
	principalInput = strings.TrimSpace(principalInput)
	principal, _ := strconv.ParseFloat(principalInput, 64)

	// Prompt for rate of interest
	fmt.Print("Enter the rate of interest: ")
	rateInput, _ := reader.ReadString('\n')
	rateInput = strings.TrimSpace(rateInput)
	rate, _ := strconv.ParseFloat(rateInput, 64)

	// Prompt for number of years
	fmt.Print("Enter the number of years: ")
	yearsInput, _ := reader.ReadString('\n')
	yearsInput = strings.TrimSpace(yearsInput)
	years, _ := strconv.ParseFloat(yearsInput, 64)

	// Compute simple interest
	rate = rate / 100 // Convert percent rate to decimal
	amount := principal * (1 + rate*years)

	// Round up to the nearest cent
	amount = float64(int(amount*100+0.999999)) / 100.0

	// Print result formatted as currency
	fmt.Printf("After %.0f years at %.1f%%, the investment will\nbe worth $%.2f.\n", years, rate*100, amount)
}
