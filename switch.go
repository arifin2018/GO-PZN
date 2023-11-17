package main

import (
	"fmt"
)

func main()  {
	var name = "arifin"

	switch name {
	case "arifin":
		fmt.Println("arifin")
	default:
		fmt.Println("else")
	}

	switch length:= len(name); length > 5 { //switch dengan kodisi
	case true:
		fmt.Println("lebih dari 5")
	case false:
		fmt.Println("kurang dari 5")
	}

	switch { //switch kosong case dengan kodisi
	case string(name[0]) == "a":
		fmt.Println("a")
	case string(name[0]) == "b":
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}