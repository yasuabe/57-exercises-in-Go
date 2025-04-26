// Exercise 11: Currency Conversion
//
// - Prompt for:
//   - Euro amount
//   - Exchange rate
// - Convert euros to U.S. dollars using:
//   - `amount_to = (amount_from × rate_from) / rate_to`
// - Round up to the next cent
// - Print result in a single output statement.

package main

import "fmt"

func main() {
	// Prompt for Euro amount and exchange rate
	var euros, exchangeRate float64

	// Get user input
	fmt.Print("How many euros are you exchanging? ")
	fmt.Scanln(&euros)
	fmt.Print("What is the exchange rate? ")
	fmt.Scanln(&exchangeRate)

	// Convert euros to U.S. dollars
	dollars := (euros * exchangeRate) / 100.0

	// Round up to the next cent
	dollars = float64(int(dollars*100+0.999999)) / 100.0

	// Print result in a single output statement
	fmt.Printf("%.2f euros at an exchange rate of %.2f is\n%.2f U.S. dollars.\n", euros, exchangeRate, dollars)
}
