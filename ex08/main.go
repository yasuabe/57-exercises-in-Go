// Ex8: Pizza Party
//
// - Prompt the user for the number of people, pizzas, and slices per pizza.
// - Calculate the total number of slices.
// - Determine how many slices each person gets evenly.
// - Calculate and display any leftover slices.
// - Output the distribution results clearly.

package main

import "fmt"

func main() {
	// Prompt the user for the number of people, pizzas, and slices per pizza
	var people, pizzas, slicesPerPizza int
	fmt.Print("How many people? ")
	fmt.Scanln(&people)
	fmt.Print("How many pizzas do you have? ")
	fmt.Scanln(&pizzas)
	fmt.Print("How many slices per pizza? ")
	fmt.Scanln(&slicesPerPizza)

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
