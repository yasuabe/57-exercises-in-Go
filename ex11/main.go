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

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for Euro amount
	fmt.Print("How many euros are you exchanging? ")
	eurosInput, _ := reader.ReadString('\n')
	eurosInput = strings.TrimSpace(eurosInput)
	euros, _ := strconv.ParseFloat(eurosInput, 64)

	// Prompt for exchange rate
	fmt.Print("What is the exchange rate? ")
	exchangeRateInput, _ := reader.ReadString('\n')
	exchangeRateInput = strings.TrimSpace(exchangeRateInput)
	exchangeRate, _ := strconv.ParseFloat(exchangeRateInput, 64)

	// Convert euros to U.S. dollars
	dollars := (euros * exchangeRate) / 100.0

	// Round up to the next cent
	dollars = float64(int(dollars*100+0.999999)) / 100.0

	// Print result in a single output statement
	fmt.Printf("%.2f euros at an exchange rate of %.2f is\n%.2f U.S. dollars.\n", euros, exchangeRate, dollars)
}
