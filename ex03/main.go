// Ex3: Printing Quotes
//
// - Prompt the user to enter a quote.
// - Prompt the user to enter the author of the quote.
// - Display the author and quote using escaped quotation marks.
// - Use string concatenation, not interpolation or substitution.
// - Use a single output statement for the result.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter a quote
	fmt.Println("What is the quote? ")
	quote, _ := reader.ReadString('\n')
	quote = quote[:len(quote)-1] // Remove the trailing newline

	// Prompt the user to enter the author of the quote
	fmt.Println("Who said it? ")
	author, _ := reader.ReadString('\n')
	author = author[:len(author)-1] // Remove the trailing newline

	// Display the author and quote using escaped quotation marks
	result := author + " says, \"" + quote + "\""

	// Use a single output statement for the result
	fmt.Println(result)
}
