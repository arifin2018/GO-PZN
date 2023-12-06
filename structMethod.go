package main

import "fmt"

type Agent struct {
	name,address	string
	age				int
}

func (agent Agent) sayHello(name string) {
	fmt.Println("hallo saya",name,"selamat datang",agent.name)
}

func main()  {
	azriel := Agent{"azriel","jakarta",10}
	azriel.sayHello("arifin") //hai saya arifin selamat datang azriel
}