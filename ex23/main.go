// Ex23: Troubleshooting Car Issues
//
// - Guide the user through a series of yes/no questions based on a decision tree.
// - Ask only relevant questions based on previous answers.
// - Display appropriate troubleshooting advice depending on the answers.
// - Do not prompt for all inputs at once; follow the flow of the decision tree.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Is the car silent when you turn the key? ")
	if getYesOrNo(reader) {
		fmt.Print("Are the battery terminals corroded? ")
		if getYesOrNo(reader) {
			fmt.Println("Clean terminals and try starting again.")
		} else {
			fmt.Println("The battery cables may be damaged.")
			fmt.Println("Replace cables and try again.")
		}
	} else {
		fmt.Print("Does the car make a clicking noise? ")
		if getYesOrNo(reader) {
			fmt.Println("Replace the battery.")
		} else {
			fmt.Print("Does the car crank up but fail to start? ")
			if getYesOrNo(reader) {
				fmt.Println("Check spark plug connections.")
			} else {
				fmt.Print("Does the engine start and then die? ")
				if getYesOrNo(reader) {
					fmt.Print("Does your car have fuel injection? ")
					if getYesOrNo(reader) {
						fmt.Println("Get it in for service.")
					} else {
						fmt.Println("Check to ensure the choke is opening and closing.")
					}
				} else {
					fmt.Println("This is a classic case of bad fuel.")
				}
			}
		}
	}
}

// getYesOrNo reads user input and returns true for "y" and false for "n".
func getYesOrNo(reader *bufio.Reader) bool {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Print("Please enter 'y' or 'n': ")
		}
	}
}
