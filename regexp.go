package main

import (
	"fmt"
	"regexp"
)

/**
regexp dari golang dibuat dari bahasa C oleh google yang diberi nama RE2
link : https://github.com/google/re2/wiki/Syntax
*/
func main(){
	regex := regexp.MustCompile(`a([a-z]+)n`)
	fmt.Println(regex.MatchString("arifin"))

	regex2 := regexp.MustCompile(`a[^rifi]+n`)
	fmt.Println(regex2.MatchString("arifin!"))
	fmt.Println(regex2.MatchString("asdxn!"))

	regex3 := regexp.MustCompile(`\d`)
	fmt.Println(regex3.MatchString("awd"))

	regex4 := regexp.MustCompile(`[a-z][\s]a[rifi]+n`)
	fmt.Println(regex4.MatchString("n arifin"))
}