package main

import "fmt"

func alphabets(data ...string) string {
	var alphabet string
	for _, v := range data {
		alphabet += v
	}
	return alphabet
}

func main(){
	alphabet := alphabets("a","b","c","d","e","f","g","h")
	fmt.Println(alphabet)
}