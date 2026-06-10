package main

import (
	"fmt"

	"github.com/obrunoapolinario/kNN/internal/csv"
)

func main() {
	customers, err := csv.ReadCustomersFromCSV("data/customers.csv")

	if err != nil {
		panic(err)
	}

	// For demonstration, print the first 5 customers
	for i, customer := range customers {
		if i >= 5 {
			break
		}
		fmt.Println(customer)
	}
}
