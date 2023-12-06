package main

import "fmt"

type Customer struct {
	Name,Address 	string
	Age				int
}

func main()  {
	var arifin Customer
	fmt.Println(arifin)

	arifin.Name = "arifin"
	arifin.Address = "jl.gandaria"
	arifin.Age = 10
	fmt.Println(arifin)

	//struct literals
	azriel := Customer{
		Name: "Azriel",
		Address: "Jagakarsa",
		Age: 30,
	}
	fmt.Println(azriel)

	lita := Customer{"Talitha","Jakarta",10} //struct literals urutan, ini harus sama dan sesuai urutan nya
	fmt.Println(lita)
}