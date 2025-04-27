//  Ex33: Magic 8 Ball
//
// - Create a Magic 8 Ball game.
// - Prompt user for a question.
// - Randomly reply with one of:
//   - “Yes”
//   - “No”
//   - “Maybe”
//   - “Ask again later”
// - Use a list (array) and a random number generator to select the response.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// List of possible Magic 8 Ball responses
	responses := []string{
		"Yes",
		"No",
		"Maybe",
		"Ask again later",
	}

	// Prompt the user for a question
	fmt.Print("What's your question? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Select a random response
	randomIndex := rng.Intn(len(responses))
	fmt.Println(responses[randomIndex])
}
