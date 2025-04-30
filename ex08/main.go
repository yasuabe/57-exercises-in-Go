// Ex8: Pizza Party
//
// - Prompt the user for the number of people, pizzas, and slices per pizza.
// - Calculate the total number of slices.
// - Determine how many slices each person gets evenly.
// - Calculate and display any leftover slices.
// - Output the distribution results clearly.

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

	// Prompt the user for the number of people
	fmt.Print("How many people? ")
	peopleInput, _ := reader.ReadString('\n')
	peopleInput = strings.TrimSpace(peopleInput)
	people, _ := strconv.Atoi(peopleInput)

	// Prompt the user for the number of pizzas
	fmt.Print("How many pizzas do you have? ")
	pizzasInput, _ := reader.ReadString('\n')
	pizzasInput = strings.TrimSpace(pizzasInput)
	pizzas, _ := strconv.Atoi(pizzasInput)

	// Prompt the user for the number of slices per pizza
	fmt.Print("How many slices per pizza? ")
	slicesPerPizzaInput, _ := reader.ReadString('\n')
	slicesPerPizzaInput = strings.TrimSpace(slicesPerPizzaInput)
	slicesPerPizza, _ := strconv.Atoi(slicesPerPizzaInput)

	// Calculate the total number of slices
	totalSlices := pizzas * slicesPerPizza

	// Determine how many slices each person gets evenly
	slicesPerPerson := totalSlices / people

	// Calculate any leftover slices
	leftoverSlices := totalSlices % people

	// Output the distribution results clearly
	fmt.Printf("%d people with %d pizzas\n", people, pizzas)
	fmt.Printf("Each person gets %d pieces of pizza.\n", slicesPerPerson)
	fmt.Printf("There are %d leftover pieces.\n", leftoverSlices)
}
