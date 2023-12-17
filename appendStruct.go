package main

import (
	"fmt"
)

type person struct{
	name string
	age int
}

type family struct{
	familys []person
} 

type setGet interface {
	setAppend() []struct{}
}

func (f *family) setAppend(dataPerson person) []person {
	f.familys = append(f.familys, dataPerson)
	return f.familys
}


func main()  {
	fams := person{name:"azriel",age: 24}
	fams2 := person{name:"azriel2",age: 24}
	items := []person{}
    familyVar := family{items}

	familyVar.setAppend(fams)
	familyVar.setAppend(fams2)
	fmt.Println(familyVar.familys)
}