package main

import "fmt"

func main()  {
	var names [2]string

	names[0] = "nur"
	names[1] = "arifin"
	fmt.Println(names)

	var names2 [2]string = [2]string{
        "nur",
        "arifin",
	}
	fmt.Println(names2)

	var names3 = [2]string{
        "nur",
        "arifin",
	}
	fmt.Println(names3)
}