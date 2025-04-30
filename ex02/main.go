// Ex2: Counting the Number of Characters
//
// - Prompt the user to enter an input string.
// - Determine the number of characters using a built-in function.
// - Output the original string and its character count.
// - Use a single output statement to construct the output.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter an input string
	fmt.Print("Enter a string: ") // Use fmt.Print to avoid a newline after the prompt
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Read all characters up to the newline

	// Trim the newline character from the input
	input = strings.TrimSpace(input)

	// Determine the number of characters using a built-in function
	charCount := len(input)

	// Output the original string and its character count
	fmt.Printf("The string '%s' has %d characters.\n", input, charCount)
}
