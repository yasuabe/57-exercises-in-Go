// # Exercise 21: Numbers to Names
//
// - Prompt the user to enter a number from 1 to 12.
// - Display the corresponding month name (e.g., 1 → January).
// - If the number is outside this range, show an error message.
// - Use a switch or case statement.
// - Use a single output statement.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Prompt the user to enter a number
	fmt.Print("Please enter the number of the month: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	monthNumber, err := strconv.Atoi(strings.TrimSpace(input))

	// Handle invalid input
	if err != nil || monthNumber < 1 || monthNumber > 12 {
		fmt.Println("Invalid input. Please enter a number between 1 and 12.")
		return
	}

	// Determine the month name using a switch statement
	monthName := func() string {
		switch monthNumber {
		case 1:
			return "January"
		case 2:
			return "February"
		case 3:
			return "March"
		case 4:
			return "April"
		case 5:
			return "May"
		case 6:
			return "June"
		case 7:
			return "July"
		case 8:
			return "August"
		case 9:
			return "September"
		case 10:
			return "October"
		case 11:
			return "November"
		case 12:
			return "December"
		default:
			return "Invalid" // This should never be reached due to earlier validation
		}
	}()

	// Output the result in a single statement
	fmt.Printf("The name of the month is %s.\n", monthName)
}
