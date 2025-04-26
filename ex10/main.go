// # Ex10: Self-Checkout
//
// - Prompt for price and quantity of 3 items.
// - Calculate subtotal, then 5.5% tax, then total.
// - Print each line item, subtotal, tax, and total.
// - Separate input, processing, and output logic.
// - Ensure all input is converted to numeric types before calculations.
package main

import "fmt"

func getItemDetails(itemNumber int) (float64, float64) {
	var price float64
	var quantity int

	fmt.Printf("Enter the price of item %d: ", itemNumber)
	fmt.Scan(&price)

	fmt.Printf("Enter the quantity of item %d: ", itemNumber)
	fmt.Scan(&quantity)

	return price, float64(quantity)
}
func calculateSubtotal(price1, quantity1, price2, quantity2, price3, quantity3 float64) float64 {
	return (price1 * float64(quantity1)) + (price2 * float64(quantity2)) + (price3 * float64(quantity3))
}
func calculateTax(subtotal float64) float64 {
	return subtotal * 0.055
}
func printReceipt(subtotal, tax, total float64) {
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Total: $%.2f\n", total)
}
func main() {
	// Prompt for price and quantity of 3 items
	price1, quantity1 := getItemDetails(1)
	price2, quantity2 := getItemDetails(2)
	price3, quantity3 := getItemDetails(3)

	// Calculate subtotal
	subtotal := calculateSubtotal(price1, quantity1, price2, quantity2, price3, quantity3)

	// Calculate tax
	tax := calculateTax(subtotal)

	// Calculate total
	total := subtotal + tax

	// Print results
	printReceipt(subtotal, tax, total)
}
