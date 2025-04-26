// Ex2: Counting the Number of Characters
//
// - Prompt the user to enter an input string.
// - Determine the number of characters using a built-in function.
// - Output the original string and its character count.
// - Use a single output statement to construct the output.
package main

import "fmt"

func main() {
	// Prompt the user to enter an input string
	var input string
	println("Enter a string:")
	fmt.Scanln(&input)

	// Determine the number of characters using a built-in function
	charCount := len(input)

	// Output the original string and its character count
	fmt.Printf("The string '%s' has %d characters.\n", input, charCount)
}
