// # Ex1: Saying Hello
//
// - Prompt the user to enter their name.
// - Create a greeting message using string concatenation.
// - Print the greeting.
// - Keep input, concatenation, and output as separate steps.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter their name
	fmt.Print("Enter your name: ") // Use fmt.Print to avoid a newline after the prompt
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // Read all characters up to the newline

	// Trim the newline character from the input
	name = strings.TrimSpace(name)

	// Create a greeting message using string concatenation
	greeting := "Hello, " + name

	// Print the greeting
	fmt.Println(greeting)
}
