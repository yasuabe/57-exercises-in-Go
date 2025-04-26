// # Ex1: Saying Hello
//
// - Prompt the user to enter their name.
// - Create a greeting message using string concatenation.
// - Print the greeting.
// - Keep input, concatenation, and output as separate steps.

package main

import "fmt"

func main() {
	// Prompt the user to enter their name
	var name string
	println("Enter your name:")
	fmt.Scanln(&name)

	// Create a greeting message using string concatenation
	greeting := "Hello, " + name

	// Print the greeting
	fmt.Println(greeting)
}
