// # Ex52: Creating Your Own Time Service
//
// - Build a minimal web server that returns the current time as JSON: { "currentTime": "2050-01-24 15:06:26" }.
// - Build a client that fetches this JSON, parses it, and displays the time in a readable format.
// - Server must set Content-Type: application/json.
// - Keep server code minimal.

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currentTime := TimeResponse{CurrentTime: time.Now().Format("2006-01-02 15:04:05")}
	json.NewEncoder(w).Encode(currentTime)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}
