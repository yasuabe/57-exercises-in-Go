// Ex32 Guess the Number Game
//
// - Guess the Number with 3 difficulty levels:
//   -  Level 1: 1–10
//   -  Level 2: 1–100
//   -  Level 3: 1–1000
// - Random number is picked based on difficulty.
// - Player guesses until correct.
// - After each guess, show “too high” or “too low”.
// - Count and display total guesses.
// - Ask to play again after winning.
// - Non-numeric input is invalid and counts as a wrong guess.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a new random generator with a unique seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Let's play Guess the Number.")

	for {
		// Select difficulty level
		fmt.Print("Pick a difficulty level (1, 2, or 3): ")
		var input string
		fmt.Scanln(&input)

		difficulty, err := strconv.Atoi(input)
		if err != nil || difficulty < 1 || difficulty > 3 {
			fmt.Println("Invalid difficulty level. Please enter 1, 2, or 3.")
			continue
		}

		// Set range based on difficulty
		var maxNumber int
		switch difficulty {
		case 1:
			maxNumber = 10
		case 2:
			maxNumber = 100
		case 3:
			maxNumber = 1000
		}

		// Generate random number using the new random generator
		target := rng.Intn(maxNumber) + 1
		fmt.Println("I have my number. What's your guess?")

		// Start guessing loop
		guessCount := 0
		for {
			fmt.Print("Your guess: ")
			fmt.Scanln(&input)

			guessCount++
			guess, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter a numeric value.")
				continue
			}

			if guess < target {
				fmt.Println("Too low. Guess again.")
			} else if guess > target {
				fmt.Println("Too high. Guess again.")
			} else {
				fmt.Printf("You got it in %d guesses!\n", guessCount)
				break
			}
		}

		// Ask to play again
		fmt.Print("Play again? (y/n): ")
		fmt.Scanln(&input)
		if strings.ToLower(input) != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
