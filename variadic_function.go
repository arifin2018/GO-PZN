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
	fmt.Println(alphabet) //abcdefgh

	dataAlphabet2 := []string{"a","b","c","d","e","f"}
	alphabet2 := alphabets(dataAlphabet2...) //abcdef
	fmt.Println(alphabet2)
}