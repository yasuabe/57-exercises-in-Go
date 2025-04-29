// # Ex45: Word Finder
//
// - Read a text file.
// - Replace every instance of "utilize" with "use".
// - Write the result to a new file.
// - Constraint: Prompt the user for the output file name.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Input file path
	inputFilePath := "./ex45/data/input.txt"

	// Read the content of the input file
	content, err := readFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Replace "utilize" with "use"
	modifiedContent := strings.ReplaceAll(content, "utilize", "use")

	// Prompt the user for the output file name
	fmt.Print("Enter the output file name: ")
	reader := bufio.NewReader(os.Stdin)
	outputFilePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading output file name:", err)
		return
	}
	outputFilePath = strings.TrimSpace(outputFilePath) // Remove any trailing newline or spaces

	// Write the modified content to the output file
	err = writeFile(outputFilePath, modifiedContent)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("File has been successfully written to", outputFilePath)
}

// readFile reads the entire content of a file and returns it as a string
func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content.String(), nil
}

// writeFile writes the given content to a file
func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
