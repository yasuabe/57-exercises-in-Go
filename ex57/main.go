// # Ex57: Trivia App
//
// - Load questions, correct answers, and wrong answers from a local file.
// - Randomize both:
//   - Question selection.
//   - Answer order (correct + distractors).
// - Ends on first incorrect answer or all correct.
// - Track number of correct answers.
// ## Constraint:
// - Use a local file (not Redis or RDB) to store the question data.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Question struct {
	Question    string   `json:"question"`
	Correct     string   `json:"correct"`
	Distractors []string `json:"distractors"`
}

func loadQuestions(filename string) ([]Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []Question
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func shuffleAnswersWithRNG(correct string, distractors []string, rng *rand.Rand) []string {
	allAnswers := append(distractors, correct)
	rng.Shuffle(len(allAnswers), func(i, j int) {
		allAnswers[i], allAnswers[j] = allAnswers[j], allAnswers[i]
	})
	return allAnswers
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a custom random number generator

	questions, err := loadQuestions("ex57/data/trivia.json")
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	correctCount := 0
	for _, q := range questions {
		fmt.Println(q.Question)
		answers := shuffleAnswersWithRNG(q.Correct, q.Distractors, rng) // Use the custom RNG

		for i, ans := range answers {
			fmt.Printf("%d. %s\n", i+1, ans)
		}

		var userChoice int
		fmt.Print("Your answer (number): ")
		_, err := fmt.Scan(&userChoice)
		if err != nil || userChoice < 1 || userChoice > len(answers) {
			fmt.Println("Invalid input. Exiting.")
			break
		}

		if answers[userChoice-1] == q.Correct {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Wrong! Game over.")
			break
		}
	}

	fmt.Printf("You answered %d questions correctly.\n", correctCount)
}
