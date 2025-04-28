// ## Ex56: Tracking Inventory
//
// - Goal: Create a program to track personal inventory.
// - Input: Item name, serial number, and estimated value.
// - Output: A tabular report in both HTML and CSV formats.
// - Constraints:
//   - Store data persistently in a local file using JSON, XML, or YAML.
//   - The value must be numeric.

package main

import (
	"encoding/csv"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const dataFilePath = "ex56/output/inventory.json" // Define the data file path as a constant

type Item struct {
	Name     string  `json:"name"`
	SerialNo string  `json:"serial_no"`
	Value    float64 `json:"value"`
}

var (
	inventory []Item
	mu        sync.Mutex
)

func main() {
	// Load inventory from file
	loadInventory()

	// Define routes
	http.HandleFunc("/ex56", inventoryHandler)
	http.HandleFunc("/ex56/csv", csvHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission
		name := r.FormValue("name")
		serialNo := r.FormValue("serial_no")
		value, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}

		// Add the new item to the inventory
		mu.Lock()
		inventory = append(inventory, Item{Name: name, SerialNo: serialNo, Value: value})
		saveInventory()
		mu.Unlock()
	}

	// Render the HTML template
	tmpl, err := template.ParseFiles("ex56/inventory.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, map[string]interface{}{
		"items": inventory,
	})
}

func csvHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=inventory.csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Serial Number", "Value"})

	// Write inventory data
	mu.Lock()
	defer mu.Unlock()
	for _, item := range inventory {
		writer.Write([]string{item.Name, item.SerialNo, strconv.FormatFloat(item.Value, 'f', 2, 64)})
	}
}

func loadInventory() {
	file, err := os.Open(dataFilePath) // Use the constant for the file path
	if err != nil {
		return // No inventory file exists
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&inventory)
}

func saveInventory() {
	file, err := os.Create(dataFilePath) // Use the constant for the file path
	if err != nil {
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(inventory)
}
