// Ex5: Simple Math
//
// - Prompt the user to enter two numbers.
// - Convert the input strings to numeric types before performing calculations.
// - Calculate the sum, difference, product, and quotient.
// - Keep input and output separate from processing logic.
// - Use a single output statement with line breaks to display the results.

package main

import "fmt"

func main() {
	// Prompt the user to enter two numbers
	var num1, num2 float64
	fmt.Print("Enter the first number? ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number? ")
	fmt.Scanln(&num2)

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
