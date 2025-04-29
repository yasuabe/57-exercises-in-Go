// # Ex36: Computing Statistics
//
// - Prompt the user to enter numbers representing response times until “done” is entered.
// - Store the numeric inputs in an array, excluding “done”.
// - Compute and display the average, minimum, maximum, and standard deviation.
// - Use loops and arrays for input and calculations.
// - Keep input, processing, and output logic separate.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Input: Collects numbers from the user until "done" is entered.
func collectNumbers() []float64 {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "done" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		numbers = append(numbers, num)
	}

	return numbers
}

// Processing: Computes average, minimum, maximum, and standard deviation.
func computeStatistics(numbers []float64) (float64, float64, float64, float64) {
	if len(numbers) == 0 {
		return 0, 0, 0, 0
	}

	sum := 0.0
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	average := sum / float64(len(numbers))

	// Calculate standard deviation
	var varianceSum float64
	for _, num := range numbers {
		varianceSum += math.Pow(num-average, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(numbers)))

	return average, min, max, stdDev
}

// Output: Displays the results.
func displayResults(numbers []float64, average, min, max, stdDev float64) {
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("The average is %.2f.\n", average)
	fmt.Printf("The minimum is %.2f.\n", min)
	fmt.Printf("The maximum is %.2f.\n", max)
	fmt.Printf("The standard deviation is %.2f.\n", stdDev)
}

func main() {
	numbers := collectNumbers()
	average, min, max, stdDev := computeStatistics(numbers)
	displayResults(numbers, average, min, max, stdDev)
}
