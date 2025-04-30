//  Ex22: Comparing Numbers
//
// - Prompt the user to enter three numbers.
// - If any numbers are the same, exit the program.
// - Otherwise, determine and display the largest number.
// - Do not use built-in functions to find the largest value.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the first number
	fmt.Print("Enter the first number: ")
	input1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	num1, err := strconv.Atoi(strings.TrimSpace(input1))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Prompt for the second number
	fmt.Print("Enter the second number: ")
	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	num2, err := strconv.Atoi(strings.TrimSpace(input2))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Prompt for the third number
	fmt.Print("Enter the third number: ")
	input3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	num3, err := strconv.Atoi(strings.TrimSpace(input3))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Check if any numbers are the same
	if num1 == num2 || num1 == num3 || num2 == num3 {
		fmt.Println("Some numbers are the same. Exiting the program.")
		return
	}

	// Determine the largest number manually
	largest := num1
	if num2 > largest {
		largest = num2
	}
	if num3 > largest {
		largest = num3
	}

	// Display the largest number
	fmt.Printf("The largest number is %d.\n", largest)
}
