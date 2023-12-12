package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFound = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return validationError
	}

	if id != "arifin" {
		return notFound
	}

	return nil
}

func main()  {
	azriel := getById("azriel")
	fmt.Println(azriel)
	arifin := getById("arifin")
	fmt.Println(arifin)

	err := getById("")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println(err)
		}else if errors.Is(err, notFound) {
			fmt.Println(err)
		}else {
			fmt.Println("unknown errors")
		}
	}
}