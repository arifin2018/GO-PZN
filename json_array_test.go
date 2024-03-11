package GoJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street  string
	Country string
}

func TestStructJson(t *testing.T) {
	customer := Customer{
		Firstname: "arifin",
		Lastname:  "azriel",
		Age:       21,
		Married:   false,
		Hobbies:   []string{"biola", "basket"},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestStructArrayJson(t *testing.T) {
	customer := Customer{
		Firstname: "nur",
		Lastname:  "Arifin",
		Age:       32,
		Married:   false,
		Hobbies:   []string{"bola", "coding"},
		Address: []Address{
			{
				Street:  "gandaria",
				Country: "jakarta",
			},
			{
				Street:  "gandaria2",
				Country: "jakarta2",
			},
		},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArray(t *testing.T) {
	data := `{"Firstname":"nur","Lastname":"Arifin","Age":32,"Married":false,"Hobbies":["bola","coding"],"Address":[{"Street":"gandaria","Country":"jakarta"},{"Street":"gandaria2","Country":"jakarta2"}]}`
	convertDataJson := map[string]any{}
	json.Unmarshal([]byte(data), &convertDataJson)
	fmt.Println(convertDataJson["Age"])
	jsonAddress, _ := json.Marshal(convertDataJson["Address"])
	fmt.Println("hallo + " + string(jsonAddress))
	mapConvertDataJsonAddress := []map[string]any{}
	err := json.Unmarshal(jsonAddress, &mapConvertDataJsonAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(mapConvertDataJsonAddress[0])
}
