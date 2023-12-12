package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var data *list.List = list.New();

	data.PushBack("a")
	// fmt.Println(data.Front())
	data.PushBack("b")
	data.PushBack("c")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}