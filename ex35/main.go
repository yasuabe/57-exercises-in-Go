// # Ex35: Picking a Winner
//
// - Prompt for names until a blank line is entered.
// - Store non-blank names in a collection.
// - Randomly select and print one name as the winner.
// - Use a loop for input and a random number generator for selection.
// - Exclude blank entries.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var names []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a name (leave blank to finish):")
	for {
		fmt.Print("Enter a name: ")
		scanner.Scan()
		name := scanner.Text()

		// Stop if the input is blank
		if name == "" {
			break
		}

		// Add non-blank names to the slice
		names = append(names, name)
	}

	// Check if there are any names entered
	if len(names) == 0 {
		fmt.Println("No names were entered.")
		return
	}

	// Create a new random generator with a time-based seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pick a random winner
	winner := names[rng.Intn(len(names))]
	fmt.Printf("The winner is... %s.\n", winner)
}
