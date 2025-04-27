// Ex40: Filtering Records
//
// - Create a program to filter employee records.
// - Search is based on first or last name containing a given substring.
// - Display matching records in a formatted table.
// - Data should be stored in an array of maps (or equivalent structure).

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Employee records stored as an array of maps
	employees := []map[string]string{
		{"FirstName": "John", "LastName": "Johnson", "Position": "Manager", "SeparationDate": "2016-12-31"},
		{"FirstName": "Tou", "LastName": "Xiong", "Position": "Software Engineer", "SeparationDate": "2016-10-05"},
		{"FirstName": "Michaela", "LastName": "Michaelson", "Position": "District Manager", "SeparationDate": "2015-12-19"},
		{"FirstName": "Jake", "LastName": "Jacobson", "Position": "Programmer", "SeparationDate": ""},
		{"FirstName": "Jacquelyn", "LastName": "Jackson", "Position": "DBA", "SeparationDate": ""},
		{"FirstName": "Sally", "LastName": "Weber", "Position": "Web Developer", "SeparationDate": "2015-12-18"},
	}

	// Prompt user for a search string
	fmt.Print("Enter a search string: ")
	var searchString string
	fmt.Scanln(&searchString)

	// Filter and display matching records
	fmt.Println("\nResults:")
	fmt.Printf("%-20s | %-18s | %-15s\n", "Name", "Position", "Separation Date")
	fmt.Println(strings.Repeat("-", 20) + "-|-" + strings.Repeat("-", 18) + "-|-" + strings.Repeat("-", 15))

	for _, employee := range employees {
		fullName := employee["FirstName"] + " " + employee["LastName"]
		if strings.Contains(strings.ToLower(fullName), strings.ToLower(searchString)) {
			fmt.Printf("%-20s | %-18s | %-15s\n", fullName, employee["Position"], employee["SeparationDate"])
		}
	}
}
