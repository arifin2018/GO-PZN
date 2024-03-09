package GoJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalObject(t *testing.T) {
	data := `{"Firstname":"Nur","Lastname":"Arifin","Age":30,"Married":false}`
	bytes := []byte(data)

	customer := &Customer{}
	err := json.Unmarshal(bytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	bytes2, _ := json.Marshal(customer)
	fmt.Println(string(bytes2))
}
