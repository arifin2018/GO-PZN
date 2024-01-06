package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main()  {
	names := []string{"nur","arifin"}
	values := []int{0,1,2,3,4}

	fmt.Println(values)
	
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Contains(names,"nur"))
	
}