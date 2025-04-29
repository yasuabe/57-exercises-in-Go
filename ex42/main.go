// # Ex42: Parsing a Data File
//
// - Read a file with comma-separated records (no CSV library).
// - Each line has: Last,First,Salary.
// - Parse lines into records manually.
// - Print a table with aligned columns using spaces.
// - Format must match the sample output.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("ex42/data/employee.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and parse the file line by line
	var records [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into fields manually
		fields := strings.Split(line, ",")
		if len(fields) == 3 {
			records = append(records, fields)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Determine column widths
	lastNameWidth := len("Last")
	firstNameWidth := len("First")
	salaryWidth := len("Salary")

	for _, record := range records {
		if len(record[0]) > lastNameWidth {
			lastNameWidth = len(record[0])
		}
		if len(record[1]) > firstNameWidth {
			firstNameWidth = len(record[1])
		}
		if len(record[2]) > salaryWidth {
			salaryWidth = len(record[2])
		}
	}

	// Print the header
	fmt.Printf("%-*s %-*s %-*s\n", lastNameWidth+1, "Last", firstNameWidth+1, "First", salaryWidth+1, "Salary")
	fmt.Println(strings.Repeat("-", lastNameWidth+firstNameWidth+salaryWidth+3))

	// Print the records
	for _, record := range records {
		fmt.Printf("%-*s %-*s %-*s\n", lastNameWidth+1, record[0], firstNameWidth+1, record[1], salaryWidth+1, record[2])
	}
}
