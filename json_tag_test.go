package GoJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Name  string `json:"Name"`
	Price int    `json:"Price"`
	QTY   int    `json:"Quantity"`
}

func TestJsonEncode(t *testing.T) {
	product := Product{
		Name:  "Xiaomi",
		Price: 2000,
		QTY:   1,
	}
	// productString := map[string]any{}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonDecode(t *testing.T) {
	data := `{"Name":"Xiaomi","Price":2000,"Quantity":1}`
	product := &Product{}
	json.Unmarshal([]byte(data), product)
	fmt.Println(product)
	fmt.Printf("%+v\n", product)
}
