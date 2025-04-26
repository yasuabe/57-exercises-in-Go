// Ex7: Area of a Rectangular Room
//
// - Prompt the user to enter the length and width of a room in feet.
// - Calculate the area in square feet.
// - Convert the area to square meters using a constant conversion factor.
// - Keep calculations separate from output.
// - Display both square feet and square meters in the output.

package main

import "fmt"

func main() {
	// Prompt the user to enter the length and width of a room in feet
	var length, width float64
	fmt.Print("What is the length of the room in feet? ")
	fmt.Scanln(&length)
	fmt.Print("What is the width of the room in feet? ")
	fmt.Scanln(&width)

	// Calculate the area in square feet
	areaFeet := length * width

	// Convert the area to square meters using a constant conversion factor
	const conversionFactor = 0.092903
	areaMeters := areaFeet * conversionFactor

	// Display the dimensions and the area in both square feet and square meters
	fmt.Printf("You entered dimensions of %.0f feet by %.0f feet.\n", length, width)
	fmt.Println("The area is")
	fmt.Printf("%.0f square feet\n", areaFeet)
	fmt.Printf("%.3f square meters\n", areaMeters)
}
