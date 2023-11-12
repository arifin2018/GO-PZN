package main

import (
	"fmt"
	"sort"
)

type sortArray struct{
	Name string
	Number int
}

func main() {
	// var input1 string = "Next Case"
	var result map[string]int;
	result= map[string]int{"orange": 5, "apple": 7,"mango": 3, "strawberry": 9}
	// apple => apple, orange, mango, strawberry
	// orange =>  apple, orange, mango, strawberry
	// strawberry => apple, orange, strawberry, mango
	// strawberry => apple, strawberry, orange, mango
	// strawberry => strawberry: 9, apple: 7 , orange: 5, mango: 3


	
	keys := make([]string, 0, len(result))
	for key,_ := range result {
        keys = append(keys, key)
    }

	sort.Slice(keys, func(i, j int) bool {
		fmt.Println("i ",i)
		fmt.Println("j ",j)
		fmt.Println("result[keys[i]] ",result[keys[i]])
		fmt.Println("result[keys[j] ",result[keys[j]])
		fmt.Println("result[keys[i]] > result[keys[j] ",result[keys[i]] > result[keys[j]])
		fmt.Println("keys ",keys)
		fmt.Println("result ",result)
		fmt.Println()
		return result[keys[i]] > result[keys[j]]
	})
	// fmt.Println("==================")
	// fmt.Println(result)
	// fmt.Println(keys)
	for _, k := range keys{
        fmt.Println(k, result[k])
    }

	// for i := 0; i < strings.Count(input1,"")-1; i++ {
	// 	if string(input1[i]) == "a" || string(input1[i]) == "i" || string(input1[i]) == "u" || string(input1[i]) == "e" || string(input1[i]) == "o" {
	// 		// result += string(input1[i])
	// 		result[string(input1[i])]++
	// 	}
	// }
}