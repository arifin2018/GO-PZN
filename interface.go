package main

import "fmt"

type name interface {
	GetName() string // nama function yang harus mengembalikan string
}

func SayHello(name name)  {
	fmt.Println("Hello", name.GetName()) //function yang memanggil interface function GetName()
}

type People struct {
	Name string
}

func (people People) GetName() string { //secara langsung function GetName sudah terikat dengan interface struct People
	return people.Name
}

func main()  {
	arifin := People{Name: "arifin"}
	SayHello(arifin)
}