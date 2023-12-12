package main

import (
	"fmt"
	"strconv"
)

func main()  {
	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(i)
	fmt.Println(s)

	u, err := strconv.ParseUint("420", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u)
	q := strconv.FormatInt(-42, 32)
	fmt.Println(q)
}