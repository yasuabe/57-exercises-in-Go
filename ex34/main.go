// # Ex34: Employee List Removal
//
// - Store a list of employee names in an array or list.
// - Display the full list of names initially.
// - Prompt the user for an employee name to remove.
// - Remove that name from the list.
// - Display the updated list, showing each remaining name on its own line.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initial list of employees
	employees := []string{
		"John Smith",
		"Jackie Jackson",
		"Chris Jones",
		"Amanda Cullen",
		"Jeremy Goodwin",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the current list of employees
		fmt.Printf("There are %d employees:\n", len(employees))
		for _, employee := range employees {
			fmt.Println(employee)
		}

		// Prompt for the name to remove
		fmt.Print("\nEnter an employee name to remove (or press Enter to quit): ")
		input, _ := reader.ReadString('\n')
		nameToRemove := strings.TrimSpace(input)

		// Exit the loop if the user presses Enter without input
		if nameToRemove == "" {
			fmt.Println("Exiting program.")
			break
		}

		// Remove the name if it exists in the list
		updatedEmployees := []string{}
		found := false
		for _, employee := range employees {
			if employee == nameToRemove {
				found = true
			} else {
				updatedEmployees = append(updatedEmployees, employee)
			}
		}

		// Update the list of employees
		if found {
			employees = updatedEmployees
		}

		// If the name was not found, the list remains unchanged
		fmt.Println()
	}
}
