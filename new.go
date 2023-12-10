package main

import "fmt"

type address struct {
	name string
}

func main() {
	address1 := new(address) 
	fmt.Println(address1)
	address2 := address1 //sama menggunakan &address1
	address2.name = "arifin"
	fmt.Println(address1) //&{arifin}
	fmt.Println(address2) //&{arifin}
}