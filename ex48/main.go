// # Ex48: Grabbing the Weather
//
// - Prompt the user for a city name.
// - Use OpenWeatherMap API to fetch current weather.
// - Display the temperature in Fahrenheit.
// - Constraint: Separate logic for fetching weather data from display logic.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Config holds the API key from the configuration file
type Config struct {
	APIKey string `json:"apiKey"`
}

// WeatherResponse represents the structure of the OpenWeatherMap API response
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

// loadConfig reads the API key from the config.json file
func loadConfig(configPath string) (Config, error) {
	var config Config
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	return config, err
}

// fetchWeather fetches weather data for a given city using the OpenWeatherMap API
func fetchWeather(apiKey, city string) (WeatherResponse, error) {
	var weather WeatherResponse
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("failed to fetch weather data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather, err
	}

	err = json.Unmarshal(body, &weather)
	return weather, err
}

// kelvinToFahrenheit converts a temperature from Kelvin to Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

// displayWeather formats and displays the weather information
func displayWeather(weather WeatherResponse) {
	tempF := kelvinToFahrenheit(weather.Main.Temp)
	fmt.Printf("%s weather:\n", weather.Name)
	fmt.Printf("%.0f degrees Fahrenheit\n", tempF)
}

func main() {
	// Load the API key from the config file
	configPath := filepath.Join("ex48", "config", "config.json")
	config, err := loadConfig(configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// Prompt the user for a city name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Where are you? ")
	city, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	city = strings.TrimSpace(city) // Remove trailing newline and spaces

	// Fetch weather data
	weather, err := fetchWeather(config.APIKey, city)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}

	// Display the weather information
	displayWeather(weather)
}
