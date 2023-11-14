package main

import "fmt"

func main()  {
	// var name string = "arifin"
	var name string = "joko"

	if name == "arifin" {
		fmt.Printf("%s \n", name)
	}else if name == "joko"{
		fmt.Printf("%s tingkir \n", name)
	}else if name == "arifin joko"{
		fmt.Printf("%s tingkir arifin \n", name)
	}else{
		fmt.Printf("%s \n", "else")
	}

	if lngthName := len(name); lngthName > 3 {
		fmt.Println("arifin")
	}else{
		fmt.Println("not arifin")
	}
}