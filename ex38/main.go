// # Ex38: Filtering Values
//
// - Prompt the user to enter a space-separated list of numbers.
// - Convert the input string into an array.
// - Use a function filterEvenNumbers(old_array) to return a new array with only even numbers.
// - Do not use built-in filter or similar features.
// - Print the filtered even numbers.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// filterEvenNumbers filters even numbers from the input array
func filterEvenNumbers(numbers []int) []int {
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}

func main() {
	// Prompt the user to enter a space-separated list of numbers
	fmt.Print("Enter a list of numbers, separated by spaces: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert the input string into an array of integers
	strNumbers := strings.Split(input, " ")
	var numbers []int
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input. Please enter only numbers separated by spaces.")
			return
		}
		numbers = append(numbers, num)
	}

	// Filter even numbers using the custom function
	evenNumbers := filterEvenNumbers(numbers)

	// Output the filtered even numbers
	fmt.Printf("The even numbers are %v.\n", evenNumbers)
}
