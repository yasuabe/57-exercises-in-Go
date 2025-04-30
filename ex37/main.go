// # Ex37: Password Generator
//
// - Prompt for:
//   - Minimum total length
//   - Number of special characters
//   - Number of digits
// - Generate a password satisfying these constraints.
// - Use character lists and randomness.
// - Constraint: Store characters in lists and ensure randomness in generation.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NUMBER_CHARS   = "0123456789"
	ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
	ALPHABET_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SPECIAL_CHARS  = "!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
	MAX_LENGTH     = 30
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for minimum length
	minLength := promptForInt(reader, "What's the minimum length? ", 2, MAX_LENGTH)

	// Prompt for number of special characters
	specialChars := promptForInt(reader, "How many special characters? ", 0, minLength)

	// Prompt for number of digits
	maxDigits := minLength - specialChars
	digits := promptForInt(reader, "How many numbers? ", 0, maxDigits)

	// Generate password
	password := generatePassword(minLength, specialChars, digits)
	fmt.Printf("Your password is\n%s\n", password)
}

// promptForInt prompts the user for an integer input within a specific range.
func promptForInt(reader *bufio.Reader, prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || value < min || value > max {
			fmt.Printf("Please enter a valid number between %d and %d.\n", min, max)
			continue
		}
		return value
	}
}

// generatePassword generates a random password based on the given constraints.
func generatePassword(minLength, specialChars, digits int) string {
	// Create a local random generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Determine the actual password length (random between minLength and MAX_LENGTH)
	passwordLength := rng.Intn(MAX_LENGTH-minLength+1) + minLength

	// Generate special characters
	password := make([]byte, 0, passwordLength)
	for i := 0; i < specialChars; i++ {
		password = append(password, SPECIAL_CHARS[rng.Intn(len(SPECIAL_CHARS))])
	}

	// Generate digits
	for i := 0; i < digits; i++ {
		password = append(password, NUMBER_CHARS[rng.Intn(len(NUMBER_CHARS))])
	}

	// Generate remaining alphabetic characters
	remaining := passwordLength - specialChars - digits
	for i := 0; i < remaining; i++ {
		if rng.Intn(2) == 0 {
			password = append(password, ALPHABET_LOWER[rng.Intn(len(ALPHABET_LOWER))])
		} else {
			password = append(password, ALPHABET_UPPER[rng.Intn(len(ALPHABET_UPPER))])
		}
	}

	// Shuffle the password
	rng.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}
