package main

import (
	"fmt"
)

// ---------- Utility Functions ----------

// Variadic function to calculate total price
func total(prices ...float64) float64 {
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	return sum
}

// Function with multiple return values
func applyDiscount(amount float64, discount float64) (float64, error) {
	if discount < 0 || discount > 100 {
		return amount, fmt.Errorf("invalid discount: %.2f%%", discount)
	}
	final := amount - (amount * discount / 100)
	return final, nil
}

// Closure for dynamic discount generator
func discountGenerator(base float64) func() float64 {
	discount := base
	return func() float64 {
		discount += 5 // increase by 5% each call
		return discount
	}
}

// ---------- Main Program ----------

func main() {
	// Available items in the store (map of item to price)
	store := map[string]float64{
		"apple":  30.0,
		"banana": 10.0,
		"milk":   25.5,
		"bread":  40.0,
	}

	// Shopping cart (slice of items)
	cart := []string{}

	// Input: adding items
	var n int
	fmt.Print("Enter number of items to buy: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var item string
		fmt.Printf("Enter item %d: ", i+1)
		fmt.Scanln(&item)

		// check if item exists in store
		if price, ok := store[item]; ok {
			cart = append(cart, item)
			fmt.Printf("Added %s (₹%.2f) to cart\n", item, price)
		} else {
			fmt.Printf("Item '%s' not available in store!\n", item)
		}
	}

	// Display cart
	fmt.Println("\n--- Shopping Cart ---")
	if len(cart) == 0 {
		fmt.Println("Cart is empty!")
		return
	}
	fmt.Println("Items:", cart)

	// Calculate total
	prices := []float64{}
	for _, item := range cart {
		prices = append(prices, store[item])
	}
	gross := total(prices...)
	fmt.Printf("Gross total: ₹%.2f\n", gross)

	// Apply discount with closure
	getNextDiscount := discountGenerator(5) // start from 5%
	discount := getNextDiscount()           // dynamic discount
	fmt.Printf("Applying %.2f%% discount...\n", discount)

	final, err := applyDiscount(gross, discount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Final payable amount: ₹%.2f\n", final)

	// Control structure: switch for payment method
	var method string
	fmt.Print("Enter payment method (cash/card/upi): ")
	fmt.Scanln(&method)

	switch method {
	case "cash":
		fmt.Println("Payment successful with Cash.")
	case "card":
		fmt.Println("Payment successful with Card.")
	case "upi":
		fmt.Println("Payment successful with UPI.")
	default:
		fmt.Println("Invalid payment method! Payment failed.")
	}
}

/*
Example run:

Enter number of items to buy: 3
Enter item 1: apple
Added apple (₹30.00) to cart
Enter item 2: milk
Added milk (₹25.50) to cart
Enter item 3: chips
Item 'chips' not available in store!

--- Shopping Cart ---
Items: [apple milk]
Gross total: ₹55.50
Applying 10.00% discount...
Final payable amount: ₹49.95
Enter payment method (cash/card/upi): upi
Payment successful with UPI.

*/
