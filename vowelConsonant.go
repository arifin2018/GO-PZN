package main

import (
	"fmt"
	"sort"
	"strings"
)


func main() {
	var input1 string = strings.ToLower("Sample Case")
	

	var vocal map[string]int
	vocal = map[string]int{}
	var consonant map[string]int
	consonant = map[string]int{}

	for i := 0; i < len(input1); i++ {
		if string(input1[i]) == "a" || string(input1[i]) == "i" || string(input1[i]) == "u" || string(input1[i]) == "e" || string(input1[i]) == "o"{
			vocal[string(input1[i])] = vocal[string(input1[i])] + 1 
		}else if string(input1[i]) != " "{
			consonant[string(input1[i])] = consonant[string(input1[i])] + 1 
		}
	}
	

	keysVocal := make([]string, 0, len(vocal))
	for k := range vocal {
		keysVocal = append(keysVocal, k)
	}

	sort.SliceStable(keysVocal, func(i, j int) bool {
		return vocal[keysVocal[i]] > vocal[keysVocal[j]]
	})

	keysConsonant := make([]string, 0, len(consonant))
	for k := range consonant {
		keysConsonant = append(keysConsonant, k)
	}

	sort.SliceStable(keysConsonant, func(i, j int) bool {
		return consonant[keysConsonant[i]] > consonant[keysConsonant[j]]
	})

	fmt.Print("vocal => ")
	for _, v := range keysVocal {
		for i := 0; i < vocal[v]; i++ {
			fmt.Printf("%s",v)
		}
		continue
	}
	fmt.Println()
	fmt.Print("consonant => ")
	for _, v := range keysConsonant {
		for i := 0; i < consonant[v]; i++ {
			fmt.Printf("%s",v)
		}
		continue
	}
}