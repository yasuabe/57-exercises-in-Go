// ## Ex47: Who’s in Space?
//
// - Access live data from the Open Notify API (http://api.open-notify.org/astros.json).
// - Parse the JSON response.
// - Display:
//     - Total number of people in space.
//     - A table of names and spacecraft.
// - Do not use pre-downloaded data.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
)

// Struct to parse the JSON response
type AstroResponse struct {
	Number int `json:"number"`
	People []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
}

func main() {
	// API URL
	const apiURL = "http://api.open-notify.org/astros.json"

	// Fetch data from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	// Decode the JSON response
	var data AstroResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Display the results
	fmt.Printf("There are %d people in space right now:\n\n", data.Number)

	// Create a tabwriter for formatted output
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "Name\t| Craft")
	fmt.Fprintln(writer, "--------------------\t|------")

	for _, person := range data.People {
		fmt.Fprintf(writer, "%s\t| %s\n", person.Name, person.Craft)
	}

	// Flush the writer to output the table
	writer.Flush()
}
