package main

import (
	"fmt"
)

func random() any {
	return 1
	return "OK"
}

func recover2()  {
	message := recover(); // mengambil nilai panic dan membuat program akan berjalan lagi
	fmt.Println("hai,",message)
}

func main() {
	result := random()
	// fmt.Println(result.(string))
	// defer recover2()
	// fmt.Println(result.(int))
	switch result.(type) {
	case string:
		fmt.Println("string", result)
	case int:
		fmt.Println("int", result)
	default:
		fmt.Println("default", result)
	}
}