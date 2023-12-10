package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFound struct {
	Message string
}

func (v *notFound) Error() string {
	return v.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "arifin" {
		return &notFound{Message: "NOT FOUND DATA"}
	}

	return nil
}

func main() {
	err := saveData("awd", nil)
	if err != nil {
		if hub,ok := err.(*validationError); ok { // ini maksudnya hub itu adalah variable yang ingin di akses di dalam if, dan variable ok mencoba conversi dari error ke validationError, kalo ok nya bernilai true maka akan masuk kedalam,ok bisa true karna berhasil konversi nilainya
			fmt.Println(hub.Message)
		}else if hub,ok := err.(*notFound); ok{
			fmt.Println(hub.Message)
		}else{
			fmt.Println(err.Error())
		}

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Message)
		case *notFound:
			fmt.Println("notFound error", finalError.Message)
		default:
			fmt.Println("error", err.Error())
		}
	} else {
		fmt.Println("sukses")
	}
}