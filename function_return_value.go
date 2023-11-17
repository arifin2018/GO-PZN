package main

import "fmt"

//sum number dengan variadic function number
func sumNumber(number ...int) int  {
	result := 0
	for _, v := range number {
		result += v
	}
	return result
}

func main()  {
	sumNumber := sumNumber(1,2,3,4,5,6,7,8,9) //ngisi data variadic
	fmt.Println(sumNumber) //45
}