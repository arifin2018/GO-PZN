package GoJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname string
	Lastname  string
	Age       int
	Married   bool
	Hobbies   []string
}

func TestCustomer(t *testing.T) {
	customer := Customer{
		Firstname: "Nur",
		Lastname:  "Arifin",
		Age:       30,
		Married:   false,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
