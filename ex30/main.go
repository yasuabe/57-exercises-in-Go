// # Ex30: Multiplication Table
//
// - Print a full multiplication table from 0×0 to 12×12.
// - Format each line as a x b = c.
// - Constraint: Use nested loops to implement the logic.

package main

import "fmt"

func main() {
	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
