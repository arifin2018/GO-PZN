package main

import (
	"fmt"
	"sort"
)

type User struct{
	name string
	age int
}

// type UserSlice []User
type UserSlice struct{
	User []User
}

func (a UserSlice) Len() int { 
	return len(a.User) 
}

func (a UserSlice) Swap(i, j int) { 
	a.User[i], a.User[j] = a.User[j], a.User[i] 
}

func (a UserSlice) Less(i, j int) bool {
	return a.User[i].age < a.User[j].age
}


func main()  {
	data := []User{
		{
			name: "Arifin",
			age:23,
		},
		{
			name: "Azrielc",
			age:24,
		},
		{
			name: "Azriels",
			age:1,
		},
		{
			name: "A",
			age:7,
		},
	}
	foo := new(UserSlice)
	foo.User = data
	sort.Sort(foo)
	fmt.Println(*foo)
}