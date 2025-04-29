// # Ex39: Sorting Records
//
// - Create a list of employee records using list of maps.
// - Each record has: First Name, Last Name, Position, Separation Date.
// - Sort records by Last Name.
// - Display in a tabular format: Name | Position | Separation Date, properly aligned.

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Employee data represented as a map
func main() {
	employees := []map[string]string{
		{"FirstName": "John", "LastName": "Johnson", "Position": "Manager", "SeparationDate": "2016-12-31"},
		{"FirstName": "Tou", "LastName": "Xiong", "Position": "Software Engineer", "SeparationDate": "2016-10-05"},
		{"FirstName": "Michaela", "LastName": "Michaelson", "Position": "District Manager", "SeparationDate": "2015-12-19"},
		{"FirstName": "Jake", "LastName": "Jacobson", "Position": "Programmer", "SeparationDate": ""},
		{"FirstName": "Jacquelyn", "LastName": "Jackson", "Position": "DBA", "SeparationDate": ""},
		{"FirstName": "Sally", "LastName": "Weber", "Position": "Web Developer", "SeparationDate": "2015-12-18"},
	}

	// Sort employees by LastName
	sort.Slice(employees, func(i, j int) bool {
		return employees[i]["LastName"] < employees[j]["LastName"]
	})

	// Print the table header
	fmt.Printf("%-20s | %-18s | %-15s\n", "Name", "Position", "Separation Date")
	fmt.Println(strings.Repeat("-", 20) + "-|-" + strings.Repeat("-", 18) + "-|-" + strings.Repeat("-", 15))

	// Print each employee record
	for _, employee := range employees {
		name := employee["FirstName"] + " " + employee["LastName"]
		position := employee["Position"]
		separationDate := employee["SeparationDate"]
		fmt.Printf("%-20s | %-18s | %-15s\n", name, position, separationDate)
	}
}
