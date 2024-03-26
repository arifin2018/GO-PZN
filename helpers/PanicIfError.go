package helpers

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
