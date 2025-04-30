// # Ex31: Karvonen Heart Rate
//
// - Prompt the user for age and resting heart rate.
// - Validate inputs to ensure both are numeric.
// - Use the Karvonen formula to calculate target heart rate.
// - Loop from 55% to 95% intensity in increments (not hard-coded).
// - Display results in a table format.

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

	// Prompt for resting heart rate
	restingHR := promptForNumber(reader, "Enter your resting heart rate: ")

	// Prompt for age
	age := promptForNumber(reader, "Enter your age: ")

	// Display the table header
	fmt.Printf("\nResting Pulse: %d\nAge: %d\n\n", restingHR, age)
	fmt.Println("Intensity    | Rate")
	fmt.Println("-------------|--------")

	// Calculate and display target heart rate for each intensity
	for intensity := 55; intensity <= 95; intensity += 5 {
		targetHeartRate := calculateTargetHeartRate(age, restingHR, float64(intensity)/100)
		fmt.Printf("%-12s | %d bpm\n", fmt.Sprintf("%d%%", intensity), targetHeartRate)
	}
}

// promptForNumber ensures valid numeric input from the user
func promptForNumber(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err == nil && number > 0 {
			return number
		}
		fmt.Println("Invalid input. Please enter a positive number.")
	}
}

// calculateTargetHeartRate computes the target heart rate using the Karvonen formula
func calculateTargetHeartRate(age, restingHR int, intensity float64) int {
	return int((float64((220-age)-restingHR) * intensity) + float64(restingHR))
}
