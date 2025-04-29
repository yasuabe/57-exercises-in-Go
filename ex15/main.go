// # Ex15: Password Validation
//
// - Prompt the user for a password.
// - Compare the input with a hardcoded known password.
// - If it matches (case-sensitive), print “Welcome!”
// - Otherwise, print “I don’t know you.”
// - Use an if/else statement for the logic.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Hardcoded password
	const correctPassword = "abc$123"

	// Prompt the user for a password
	fmt.Print("What is the password? ")

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim any trailing newline or spaces
	input = strings.TrimSpace(input)

	// Check if the input matches the hardcoded password
	if input == correctPassword {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("I don't know you.")
	}
}
