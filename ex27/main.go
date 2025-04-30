// Ex27: Validating Inputs
//
// - Prompt the user for first name, last name, ZIP code, and employee ID.
// - Validate each input:
//   - First and last names must not be empty and must be at least two characters long.
//   - ZIP code must be numeric.
//   - Employee ID must match the format AA-1234.
//
// - Create a separate function for each validation.
// - Create a validateInput function to coordinate all validations.
// - Use a single output statement to display all error messages or success.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// validateFirstName checks if the first name is at least two characters long.
func validateFirstName(firstName string) string {
	if len(firstName) < 2 {
		return fmt.Sprintf("\"%s\" is not a valid first name. It is too short.", firstName)
	}
	return ""
}

// validateLastName checks if the last name is not empty.
func validateLastName(lastName string) string {
	if strings.TrimSpace(lastName) == "" {
		return "The last name must be filled in."
	}
	return ""
}

// validateZIPCode checks if the ZIP code is numeric.
func validateZIPCode(zipCode string) string {
	if _, err := strconv.Atoi(zipCode); err != nil {
		return "The ZIP code must be numeric."
	}
	return ""
}

// validateEmployeeID checks if the employee ID matches the format AA-1234.
func validateEmployeeID(employeeID string) string {
	matched, _ := regexp.MatchString(`^[A-Z]{2}-\d{4}$`, employeeID)
	if !matched {
		return fmt.Sprintf("%s is not a valid ID.", employeeID)
	}
	return ""
}

// validateInput coordinates all validations and returns error messages or success.
func validateInput(firstName, lastName, zipCode, employeeID string) string {
	var errors []string

	if err := validateFirstName(firstName); err != "" {
		errors = append(errors, err)
	}
	if err := validateLastName(lastName); err != "" {
		errors = append(errors, err)
	}
	if err := validateZIPCode(zipCode); err != "" {
		errors = append(errors, err)
	}
	if err := validateEmployeeID(employeeID); err != "" {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return strings.Join(errors, "\n")
	}
	return "There were no errors found."
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter the last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter the ZIP code: ")
	zipCode, _ := reader.ReadString('\n')
	zipCode = strings.TrimSpace(zipCode)

	fmt.Print("Enter an employee ID: ")
	employeeID, _ := reader.ReadString('\n')
	employeeID = strings.TrimSpace(employeeID)

	result := validateInput(firstName, lastName, zipCode, employeeID)
	fmt.Println(result)
}
