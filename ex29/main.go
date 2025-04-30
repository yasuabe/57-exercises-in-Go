// # 29: Handling Bad Input
//
// - Prompt the user for the rate of return.
// - Keep prompting until a valid, non-zero numeric value is entered.
// - Use the formula years = 72 / r to calculate the years to double the investment.
// - Display the result after receiving valid input.
// - Use a loop to handle invalid input without exiting the program.

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

	for {
		fmt.Print("What is the rate of return? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sorry. That's not a valid input.")
			continue
		}

		// Trim whitespace and newline characters
		input = strings.TrimSpace(input)

		// Convert input to a float
		rate, err := strconv.ParseFloat(input, 64)
		if err != nil || rate <= 0 {
			fmt.Println("Sorry. That's not a valid input.")
			continue
		}

		// Calculate years to double the investment
		years := 72 / rate
		fmt.Printf("It will take %.0f years to double your initial investment.\n", years)
		break
	}
}
