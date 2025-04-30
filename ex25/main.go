// # Ex 25: Password Strength Indicator
//
// - Classify a password as very weak, weak, strong, or very strong based on character type and length.
// - Implement a `passwordValidator` function that returns an enum or code (not a string).
// - Constraint: Output the result with a single print statement.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Enum-like constants for password strength
const (
	VeryWeak = iota
	Weak
	Strong
	VeryStrong
)

// passwordValidator determines the strength of a password
func passwordValidator(password string) int {
	hasLetters := false
	hasDigits := false
	hasSpecial := false
	longEnough := len(password) >= 8

	// Check character types
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetters = true
		case unicode.IsDigit(char):
			hasDigits = true
		case strings.ContainsRune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", char):
			hasSpecial = true
		}
	}

	// Handle invalid input (no valid characters)
	if !hasLetters && !hasDigits && !hasSpecial {
		return -1 // Special value for invalid input
	}

	// Match conditions
	switch {
	case longEnough && hasDigits && hasLetters && hasSpecial:
		return VeryStrong
	case longEnough:
		return Strong
	case hasDigits && !hasLetters && !hasSpecial:
		return VeryWeak
	default:
		return Weak
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a password: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}

	// Trim whitespace and newline characters
	password := strings.TrimSpace(input)

	// Validate password strength
	strength := passwordValidator(password)

	// Output the result with a single print statement
	switch strength {
	case -1:
		fmt.Printf("The password '%s' is invalid.\n", password)
	case VeryWeak:
		fmt.Printf("The password '%s' is a very weak password.\n", password)
	case Weak:
		fmt.Printf("The password '%s' is a weak password.\n", password)
	case Strong:
		fmt.Printf("The password '%s' is a strong password.\n", password)
	case VeryStrong:
		fmt.Printf("The password '%s' is a very strong password.\n", password)
	}
}
