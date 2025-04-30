// Ex13: Determining Compound Interest
//
// - Prompt the user for principal amount, interest rate (as a percentage), number of years, and compounding frequency per year.
// - Convert the interest rate by dividing it by 100.
// - Use the compound interest formula to compute the final amount.
// - Round up fractions of a cent to the next penny.
// - Format the output as money.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for principal
	fmt.Print("What is the principal amount? ")
	principalInput, _ := reader.ReadString('\n')
	principalInput = strings.TrimSpace(principalInput)
	principal, _ := strconv.ParseFloat(principalInput, 64)

	// Prompt for rate
	fmt.Print("What is the rate? ")
	rateInput, _ := reader.ReadString('\n')
	rateInput = strings.TrimSpace(rateInput)
	rate, _ := strconv.ParseFloat(rateInput, 64)

	// Prompt for number of years
	fmt.Print("What is the number of years? ")
	yearsInput, _ := reader.ReadString('\n')
	yearsInput = strings.TrimSpace(yearsInput)
	years, _ := strconv.Atoi(yearsInput)

	// Prompt for compounding frequency
	fmt.Print("What is the number of times the interest is compounded per year? ")
	timesInput, _ := reader.ReadString('\n')
	timesInput = strings.TrimSpace(timesInput)
	times, _ := strconv.Atoi(timesInput)

	// Convert percent rate to decimal
	rate = rate / 100

	// Compute compound interest using the correct formula
	amount := principal * math.Pow(1+rate/float64(times), float64(times*years))

	// Round up to the nearest cent
	amount = float64(int(amount*100+0.999999)) / 100.0

	// Print result formatted as currency
	fmt.Printf(
		"$%.2f invested at %.1f%% for %d years compounded %d times per year is $%.2f.\n",
		principal,
		rate*100,
		years,
		times,
		amount)
}
