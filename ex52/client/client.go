package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/time")
	if err != nil {
		fmt.Println("Error fetching time:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server returned non-OK status:", resp.Status)
		os.Exit(1)
	}

	var timeResponse TimeResponse
	if err := json.NewDecoder(resp.Body).Decode(&timeResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Current Time:", timeResponse.CurrentTime)
}
