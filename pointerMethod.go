package main

import "fmt"

type address struct {
	name string
}

func (address *address) Mr(name string) {
	address.name = name
}

func main()  {
	arifin := &address{}
	arifin.Mr("arifin")
	fmt.Println(arifin)
	
	azriel := address{}
	azriel.Mr("azriel")
	fmt.Println(azriel)
	fmt.Println(arifin)
}