// Ex4: Mad Lib
// - Prompt the user to enter a noun, a verb, an adjective, and an adverb.
// - Create a story using the inputs.
// - Use string interpolation or substitution to build the output.
// - Use a single output statement to display the story.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Helper function to prompt and read input
func promptInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	// Use the helper function to get inputs
	noun := promptInput("Enter a noun: ")
	verb := promptInput("Enter a verb: ")
	adjective := promptInput("Enter an adjective: ")
	adverb := promptInput("Enter an adverb: ")

	// Create a story using the inputs
	story := fmt.Sprintf("Do you %s your %s %s %s? That's hilarious!", verb, adjective, noun, adverb)

	// Use a single output statement to display the story
	fmt.Println(story)
}
