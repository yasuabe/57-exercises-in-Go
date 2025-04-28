// # Ex51: Pushing Notes to Firebase
//
// - Build a command-line app that can:
//   - Save a note: mynotes new <text>
//   - Show all notes: mynotes show
//
// - Notes are saved to Firebase using its REST API (not client libraries).
// - Notes should be stored with a timestamp and displayed in reverse chronological order.
// ### Constraints:
// - Use a config file to store the Firebase API key (not hardcoded).
// - Communicate via raw HTTP requests to Firebase's REST endpoint.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	configPath    = "ex51/config/ex51_config.json"
	idTokenPath   = "ex51/output/id_token.json"
	firebaseNotes = "/notes.json"
)

type Config struct {
	ProjectID string `json:"projectId"`
	Region    string `json:"region"`
	APIKey    string `json:"apiKey"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type IDTokenInfo struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

type Note struct {
	Date string `json:"date"`
	Note string `json:"note"`
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: mynotes <command> [arguments]")
		return
	}

	command := args[0]
	switch command {
	case "new":
		if len(args) < 2 {
			fmt.Println("Usage: mynotes new <text>")
			return
		}
		text := strings.Join(args[1:], " ") // Join all arguments after "new" into a single string
		err := saveNewNote(text)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Your note was saved.")
		}
	case "show":
		err := showNotes()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	default:
		fmt.Println("Unknown command. Use 'new' or 'show'.")
	}
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile(configPath) // Replace ioutil.ReadFile with os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &config, nil
}

func loadIDToken() (*IDTokenInfo, error) {
	data, err := os.ReadFile(idTokenPath) // Replace ioutil.ReadFile with os.ReadFile
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil // No token file exists
		}
		return nil, fmt.Errorf("failed to read ID token file: %w", err)
	}
	var tokenInfo IDTokenInfo
	err = json.Unmarshal(data, &tokenInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ID token file: %w", err)
	}
	return &tokenInfo, nil
}

func saveIDToken(tokenInfo *IDTokenInfo) error {
	data, err := json.MarshalIndent(tokenInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize ID token: %w", err)
	}
	err = os.WriteFile(idTokenPath, data, 0644) // Replace ioutil.WriteFile with os.WriteFile
	if err != nil {
		return fmt.Errorf("failed to save ID token file: %w", err)
	}
	return nil
}

func authenticate(config *Config) (*IDTokenInfo, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", config.APIKey)
	payload := map[string]string{
		"email":             config.Email,
		"password":          config.Password,
		"returnSecureToken": "true",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize authentication payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("authentication request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication failed with status: %s", resp.Status)
	}

	var res struct {
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
	}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse authentication response: %w", err)
	}

	expiresIn, err := time.ParseDuration(res.ExpiresIn + "s")
	if err != nil {
		return nil, fmt.Errorf("failed to parse token expiration: %w", err)
	}

	tokenInfo := &IDTokenInfo{
		IDToken:      res.IDToken,
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(expiresIn).Unix(),
	}
	err = saveIDToken(tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func getIDToken(config *Config) (string, error) {
	tokenInfo, err := loadIDToken()
	if err != nil {
		return "", err
	}

	if tokenInfo == nil || time.Now().Unix() >= tokenInfo.ExpiresAt {
		tokenInfo, err = authenticate(config)
		if err != nil {
			return "", err
		}
	}
	return tokenInfo.IDToken, nil
}

func getFirebaseURL(config *Config) string {
	if config.Region == "us-central1" {
		return fmt.Sprintf("https://%s-default-rtdb.firebaseio.com%s", config.ProjectID, firebaseNotes)
	}
	return fmt.Sprintf("https://%s-default-rtdb.%s.firebasedatabase.app%s", config.ProjectID, config.Region, firebaseNotes)
}

func saveNewNote(noteText string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	idToken, err := getIDToken(config)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?auth=%s", getFirebaseURL(config), idToken)
	note := Note{
		Date: time.Now().Format("2006-01-02"),
		Note: noteText,
	}
	data, err := json.Marshal(note)
	if err != nil {
		return fmt.Errorf("failed to serialize note: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to save note: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to save note with status: %s", resp.Status)
	}
	return nil
}

func showNotes() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	idToken, err := getIDToken(config)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?auth=%s", getFirebaseURL(config), idToken)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch notes: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch notes with status: %s", resp.Status)
	}

	var notesMap map[string]Note
	err = json.NewDecoder(resp.Body).Decode(&notesMap)
	if err != nil {
		return fmt.Errorf("failed to parse notes: %w", err)
	}

	var notes []Note
	for _, note := range notesMap {
		notes = append(notes, note)
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Date > notes[j].Date
	})

	for _, note := range notes {
		fmt.Printf("%s - %s\n", note.Date, note.Note)
	}
	return nil
}
