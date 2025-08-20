package main

import (
	"fmt"
)

var orders = map[string]float64{
	"burger": 5.99,
	"fries": 2.49,
	"drink": 1.99,
}

type Bill struct {
	name string
	items map[string]float64
	tip float64
}

func main() {
	for k,v := range orders {
		fmt.Println("Item:", k, "Price:", v)
	}
}