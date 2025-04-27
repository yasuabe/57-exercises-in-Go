// # Ex46: Word Frequency Finder
//
// - Read a text file.
// - Count word frequencies.
// - Display a histogram using * to show counts.
// - Sort output from most frequent to least frequent.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Input file path
	inputFilePath := "./ex46/data/words.txt"

	// Read words from the file
	words, err := readWordsFromFile(inputFilePath)
	if err != nil {
		return
	}

	// Count word frequencies
	wordCounts := countWordFrequencies(words)

	// Sort words by frequency (descending) and alphabetically for ties
	sortedWords := sortWordsByFrequency(wordCounts)

	// Display the histogram
	writeHistogram(sortedWords, wordCounts)
}

// readWordsFromFile reads words from a file and returns them as a slice of strings
func readWordsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Split input into words
	for scanner.Scan() {
		words = append(words, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// countWordFrequencies counts the frequency of each word in the slice
func countWordFrequencies(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

// sortWordsByFrequency sorts words by frequency (descending) and alphabetically for ties
func sortWordsByFrequency(wordCounts map[string]int) []string {
	words := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if wordCounts[words[i]] == wordCounts[words[j]] {
			return words[i] < words[j] // Alphabetical order for ties
		}
		return wordCounts[words[i]] > wordCounts[words[j]] // Descending frequency
	})

	return words
}

// writeHistogram displays the word frequency histogram
func writeHistogram(sortedWords []string, wordCounts map[string]int) {
	for _, word := range sortedWords {
		count := wordCounts[word]
		histogram := strings.Repeat("*", count)
		fmt.Printf("%-10s %s\n", word+":", histogram)
	}
}
