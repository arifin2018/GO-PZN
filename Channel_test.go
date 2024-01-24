package gopzn

import (
	"fmt"
	"testing"
	"time"
)

func Name(name string) string {
	return name
}

func Channel()  {
	channel := make(chan string)
	channel <- "arifin"
	close(channel)
}

func TestChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)
	
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "arifin"
		fmt.Println("hai")
	}()

	data := <- channel

	fmt.Println(data)
}

type profile struct {
	name string
	age int32
}

func ChannelParams(channel chan interface{})  {
	time.Sleep(2 * time.Second)
	channel <- &profile{
		name: "arifin",
		age: 23,
	}
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan any)
	defer close(channel)

	go ChannelParams(channel)

	data := <- channel
	fmt.Println(data.(*profile).name)
}