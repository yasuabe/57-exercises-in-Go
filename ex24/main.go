// Ex24: Anagram Checker
//
// - Prompt the user to enter two strings.
// - Check if the strings are anagrams.
// - Use a function called isAnagram that takes two strings and returns true or false.
// - Ensure both strings are the same length before checking further.
// - Display whether the two strings are anagrams.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter two strings and I'll tell you if they are anagrams:")
	fmt.Print("Enter the first string: ")
	firstString := readInput(reader)
	fmt.Print("Enter the second string: ")
	secondString := readInput(reader)

	if isAnagram(firstString, secondString) {
		fmt.Printf("\"%s\" and \"%s\" are anagrams.\n", firstString, secondString)
	} else {
		fmt.Printf("\"%s\" and \"%s\" are not anagrams.\n", firstString, secondString)
	}
}

// isAnagram checks if two strings are anagrams.
func isAnagram(str1, str2 string) bool {
	// Check if lengths are the same
	if len(str1) != len(str2) {
		return false
	}

	// Normalize strings (convert to lowercase)
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Sort the characters in both strings
	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	// Compare sorted strings
	return sortedStr1 == sortedStr2
}

// sortString sorts the characters of a string alphabetically.
func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes) // Modernized sorting using slices.Sort
	return string(runes)
}

// readInput reads a line of input from the user and trims whitespace.
func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
