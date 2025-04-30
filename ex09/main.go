// Ex9: Paint Calculator
//
// - Prompt for room length and width.
// - Compute area and divide by 350 (coverage per gallon).
// - Round up to the next whole gallon.
// - Use a constant for coverage rate.

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
	const coveragePerGallon = 350.0 // coverage per gallon in square feet

	// Prompt for room length
	fmt.Print("Enter the length of the room in feet: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, _ := strconv.ParseFloat(lengthInput, 64)

	// Prompt for room width
	fmt.Print("Enter the width of the room in feet: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)
	width, _ := strconv.ParseFloat(widthInput, 64)

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
