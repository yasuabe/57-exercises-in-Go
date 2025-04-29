// # Ex 26: Months to Pay Off a Credit Card
//
// - Prompt the user for credit card balance, APR (as a percentage), and monthly payment.
// - Calculate how many months are needed to pay off the balance using the given formula.
// - Internally convert APR to a daily rate.
// - Use a function calculateMonthsUntilPaidOff(balance, apr, payment) to perform the calculation.
// - Round up any fractional result to the next whole number.
// - Do not access input values outside the function.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// calculateMonthsUntilPaidOff calculates the number of months to pay off a credit card balance.
func calculateMonthsUntilPaidOff(balance, apr, payment float64) int {
	dailyRate := apr / 365 / 100 // Convert APR to daily rate
	numerator := math.Log(1 + balance/payment*(1-math.Pow(1+dailyRate, 30)))
	denominator := math.Log(1 + dailyRate)
	months := -(1.0 / 30.0) * numerator / denominator
	return int(math.Ceil(months)) // Round up to the next whole number
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for balance
	fmt.Print("What is your balance? ")
	balanceInput, _ := reader.ReadString('\n')
	balance, err := strconv.ParseFloat(strings.TrimSpace(balanceInput), 64)
	if err != nil {
		fmt.Println("Invalid input for balance.")
		return
	}

	// Prompt for APR
	fmt.Print("What is the APR on the card (as a percent)? ")
	aprInput, _ := reader.ReadString('\n')
	apr, err := strconv.ParseFloat(strings.TrimSpace(aprInput), 64)
	if err != nil {
		fmt.Println("Invalid input for APR.")
		return
	}

	// Prompt for monthly payment
	fmt.Print("What is the monthly payment you can make? ")
	paymentInput, _ := reader.ReadString('\n')
	payment, err := strconv.ParseFloat(strings.TrimSpace(paymentInput), 64)
	if err != nil {
		fmt.Println("Invalid input for monthly payment.")
		return
	}

	// Calculate months to pay off
	months := calculateMonthsUntilPaidOff(balance, apr, payment)

	// Output result
	fmt.Printf("\nIt will take you %d months to pay off this card.\n", months)
}
