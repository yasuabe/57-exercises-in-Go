// Ex9: Paint Calculator
//
// - Prompt for room length and width.
// - Compute area and divide by 350 (coverage per gallon).
// - Round up to the next whole gallon.
// - Use a constant for coverage rate.

package main

import "fmt"

func main() {
	// Prompt for room length and width
	var length, width float64
	const coveragePerGallon = 350.0 // coverage per gallon in square feet

	// Get user input
	fmt.Print("Enter the length of the room in feet: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width of the room in feet: ")
	fmt.Scanln(&width)

	// Calculate area and gallons needed
	area := length * width
	gallonsNeeded := area / coveragePerGallon

	// Round up to the next whole gallon
	if gallonsNeeded != float64(int(gallonsNeeded)) {
		gallonsNeeded = float64(int(gallonsNeeded) + 1)
	}

	// Output the result
	fmt.Printf("You will need to purchase %.0f gallons of paint to cover %.0f square feet.\n", gallonsNeeded, area)
}
