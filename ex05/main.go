// Ex5: Simple Math
//
// - Prompt the user to enter two numbers.
// - Convert the input strings to numeric types before performing calculations.
// - Calculate the sum, difference, product, and quotient.
// - Keep input and output separate from processing logic.
// - Use a single output statement with line breaks to display the results.

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

	// Prompt the user to enter the first number
	fmt.Print("Enter the first number? ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, _ := strconv.ParseFloat(input1, 64)

	// Prompt the user to enter the second number
	fmt.Print("Enter the second number? ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, _ := strconv.ParseFloat(input2, 64)

	// Calculate the sum, difference, product, and quotient
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Use a single output statement with line breaks to display the results
	fmt.Printf("%.0f + %.0f = %.0f\n", num1, num2, sum)
	fmt.Printf("%.0f - %.0f = %.0f\n", num1, num2, difference)
	fmt.Printf("%.0f * %.0f = %.0f\n", num1, num2, product)
	fmt.Printf("%.0f / %.0f = %.0f\n", num1, num2, quotient)
}
