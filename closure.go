package main

import "fmt"

func main(){
	counter := 0
	
	closure := func ()  {
		fmt.Println("counter")
		counter++
		counter++
	}

	closure()
	fmt.Print(counter)
}