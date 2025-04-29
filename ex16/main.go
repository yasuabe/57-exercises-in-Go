// # Ex16: Legal Driving Age
//
// - Prompt the user for their age.
// - Compare it to the legal driving age (16).
// - Output a single message:
//   - If 16 or older → “You are old enough to legally drive.”
//   - If under 16 → “You are not old enough to legally drive.”
// - Use a ternary operator if available, or if/else with a single print statement.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const legalDrivingAge = 16

	// Prompt the user for their age
	fmt.Print("What is your age? ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(input))

	// Handle invalid input
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Output a single message using an inline if/else
	fmt.Println(
		func() string {
			if age >= legalDrivingAge {
				return "You are old enough to legally drive."
			}
			return "You are not old enough to legally drive."
		}(),
	)
}
