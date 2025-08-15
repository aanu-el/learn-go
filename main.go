package main

import (
	"fmt"
)

func main() {
	
	menu := map[string]float64{
		"Burger": 5.99,
		"Fries": 2.49,
		"Drink": 1.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["Burger"])

	phonebook := map[int]string{
		1234567890: "Alice",
		9876543210: "Bob",
		5555555555: "Charlie",
	}

	for k, v := range phonebook {
		fmt.Printf("Phone: %d, Name: %s\n", k, v)
	}
}