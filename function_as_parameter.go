package main

import "fmt"

type filtered func(string) string //type function decleration

func sayHello(name string, filter filtered) {// supaya tidak kepanjangan makanya dibuat filtered
	fmt.Println(filter(name))
}

func sayPrint(name string, print func(name string) string) {
	fmt.Println(print(name))
}

func filter(name string) string {
	if name == "kunti" {
		return "setan"
	}else{
		return name
	}
}

func main()  {
	sayPrint("arifin", filter)
	sayPrint("kunti", filter)
}