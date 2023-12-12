package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("nur arifin", "arifin"))
	fmt.Println(strings.Split("nur arifin", " "))
	fmt.Println(strings.ToLower("nur arifin"))
	fmt.Println(strings.ToUpper("nur arifin"))
	fmt.Println(strings.Trim("   nur arifin   ","    "))
	fmt.Println(strings.Replace("nur arifin nur arifin","nur","aksen",1))
	fmt.Println(strings.ReplaceAll("nur arifin nur arifin","nur","aksen"))

}