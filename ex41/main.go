// Ex41: Name Sorter
//
// - Read a list of names from a file.
// - Sort the names alphabetically.
// - Output:
//    - Total number of names.
//    - A separator line.
//    - The sorted names.
// - Do not hard-code the number of names.

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Input and output file paths
	inputFilePath := "./ex41/data/names.txt"
	outputFilePath := "./ex41/output/names_sorted.txt"

	// Read names from the input file
	names, err := readNamesFromFile(inputFilePath)
	if err != nil {
		return
	}

	// Sort names alphabetically
	sort.Strings(names)

	// Write sorted names to the output file
	_ = writeNamesToFile(outputFilePath, names)
}

// readNamesFromFile reads names from a file and returns them as a slice of strings
func readNamesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return names, nil
}

// writeNamesToFile writes a slice of names to a file
func writeNamesToFile(filePath string, names []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write the total number of names and a separator line
	_, err = writer.WriteString("Total of " + strconv.Itoa(len(names)) + " names\n")
	if err != nil {
		return err
	}
	_, err = writer.WriteString(strings.Repeat("-", 20) + "\n")
	if err != nil {
		return err
	}

	// Write the sorted names
	for _, name := range names {
		_, err := writer.WriteString(name + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
