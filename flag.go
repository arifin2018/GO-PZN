package main

import (
	"flag"
	"fmt"
)

func main()  {
	var username = flag.String("username","root","database username")
	var password = flag.String("password","root","database password")
	var host = flag.String("host","root","database host")
	var port = flag.String("port","root","database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", password)
	fmt.Println("host", host)
	fmt.Println("port", port)
	// go run flag.go -username="arifin"
}