// # Ex44: Product Search
//
// - Prompt user for a product name.
// - Load product data from a JSON file.
// - Search for a matching product.
// - If found, display its name, price, and quantity.
// - If not found, prompt again.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Product represents the structure of a product in the JSON file.
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// ProductData represents the top-level structure of the JSON file.
type ProductData struct {
	Products []Product `json:"products"`
}

func main() {
	// Load product data from JSON file
	products, err := loadProducts("ex44/data/products.json")
	if err != nil {
		fmt.Println("Error loading product data:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt user for a product name
		fmt.Print("What is the product name? ")

		// Read a single line of input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim whitespace and newline characters
		productName := strings.TrimSpace(input)

		// Search for the product
		product, found := searchProduct(products, productName)
		if found {
			// Display product details
			fmt.Printf("Name: %s\nPrice: $%.2f\nQuantity on hand: %d\n", product.Name, product.Price, product.Quantity)
			break
		} else {
			fmt.Println("Sorry, that product was not found in our inventory.")
		}
	}
}

// loadProducts loads product data from a JSON file.
func loadProducts(filename string) ([]Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var productData ProductData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&productData); err != nil {
		return nil, err
	}

	return productData.Products, nil
}

// searchProduct searches for a product by name in the product list.
func searchProduct(products []Product, name string) (Product, bool) {
	for _, product := range products {
		if strings.EqualFold(product.Name, name) {
			return product, true
		}
	}
	return Product{}, false
}
