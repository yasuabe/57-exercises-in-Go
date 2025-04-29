// # Ex50: Movie Recommendations
//
// - Prompt for a movie title.
// - Fetch and display: title, year, rating, runtime, synopsis.
// - Recommend based on Rotten Tomatoes audience score:
//   - ≥80% → recommend watching
//   - <50% → recommend avoiding
// - Use Rotten Tomatoes API with an API key

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	Title   string `json:"Title"`
	Year    string `json:"Year"`
	Rated   string `json:"Rated"`
	Runtime string `json:"Runtime"`
	Plot    string `json:"Plot"`
	Ratings []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
}

func main() {
	// Prompt for movie title
	fmt.Print("Enter the name of a movie: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	title := scanner.Text()
	title = strings.ReplaceAll(title, " ", "+")

	// Get API key from environment variable
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OMDB_API_KEY environment variable not set.")
		return
	}

	// Build the request URL
	url := fmt.Sprintf("https://www.omdbapi.com/?t=%s&apikey=%s", title, apiKey)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching movie details:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Check if the movie was found
	if movie.Title == "" {
		fmt.Println("Movie not found.")
		return
	}

	// Display movie details
	fmt.Printf("\nTitle: %s\nYear: %s\nRating: %s\nRunning Time: %s\n\nDescription: %s\n\n",
		movie.Title, movie.Year, movie.Rated, movie.Runtime, movie.Plot)

	// Find Rotten Tomatoes rating
	var rtScore string
	for _, rating := range movie.Ratings {
		if rating.Source == "Rotten Tomatoes" {
			rtScore = rating.Value
			break
		}
	}

	// Recommend based on Rotten Tomatoes score
	if rtScore != "" {
		fmt.Printf("Rotten Tomatoes Score: %s\n", rtScore)
		if strings.HasSuffix(rtScore, "%") {
			score := strings.TrimSuffix(rtScore, "%")
			if s, err := strconv.Atoi(score); err == nil {
				if s >= 80 {
					fmt.Println("You should watch this movie right now!")
				} else if s < 50 {
					fmt.Println("You should avoid this movie.")
				} else {
					fmt.Println("This movie might be worth watching.")
				}
			}
		}
	} else {
		fmt.Println("Rotten Tomatoes score not available.")
	}
}
