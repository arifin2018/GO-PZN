package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"11"`
}

type Person struct {
	Name string
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name :",valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "is", valueField.Type)
		_,ok:=valueField.Tag.Lookup("max")
		if ok {
			fmt.Println(valueType.Name() , "exist max tag")
		}else{
			fmt.Println(valueType.Name() , "doesn't exist max tag")
		}
	}
}

func main() {
	readField(Sample{Name: "arifin"})
	readField(Person{Name: "arifin"})
}