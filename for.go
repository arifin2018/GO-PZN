package main

import "fmt"

func main()  {
	// counter := 1;
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }
	
	// for dengan statment
	for i := 1; i <10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	j := 9
	k := 9
	for i := 1; i < 10; i++ {
		for j > 1 {
			fmt.Print(" ")
			j--
		}
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		k -= 1
		j = k
		fmt.Println()
	}

	//forr range
	data := []string{"arifin","azriel","MJ"}
	for _, v := range data {
		fmt.Println(v)
	}

}