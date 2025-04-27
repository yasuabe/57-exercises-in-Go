// # Ex28: Adding Numbers
//
// - Prompt the user to enter five numbers.
// - Use a counted loop to handle repeated prompting.
// - Compute the total of the entered numbers.
// - Display the total at the end.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	const totalNumbers = 5
	var total int

	for i := 1; i <= totalNumbers; i++ {
		var input string
		fmt.Printf("Enter a number: ")
		fmt.Scanln(&input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			i-- // Decrement the counter to retry this iteration
			continue
		}

		total += number
	}
	fmt.Printf("The total is %d.\n", total)
}
