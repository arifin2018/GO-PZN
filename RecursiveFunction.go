package main

import (
	"fmt"
)


func loopFunc(number int) int {
	result := 1

	for i := number; i >= 1; i-- {
		result *= i
	}

	return result
}


func recursiveFunc(number int) int {
	if number == 1 {
		return number
	}
	return number * recursiveFunc(number-1)
}

var data = [...]int{0,1,2,3,4,5,6}
var number int
var number2 = 0 //1,
var resultRecursiveArray = false

func recursiveArray(num int) int {
	if num == data[number2]{
		return -1
	}else if num != data[number2] {
		if number2 == (len(data)-1) {
			return num
		}
	} 
	number2 = number2 +1
	return recursiveArray(num)
}

func main()  {
	number := 1*2*3*4*5*6*7*8*9;
	fmt.Println(number) //362880
	fmt.Println(recursiveFunc(9)) //362880
	fmt.Println(recursiveArray(6)) //-1
	
}