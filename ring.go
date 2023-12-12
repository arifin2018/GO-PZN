package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main()  {
	var data = ring.New(9)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// data.Value = "value 2"

	// data = data.Next()
	// data.Value = "value 3"

	// data = data.Next()
	// data.Value = "value 4"

	data.Do(func(a any) {
		fmt.Println(a)
	})
}