package gopzn

import (
	"fmt"
	"strconv"
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

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	var data string

	go OnlyIn(channel)
	OnlyOut(channel, &data)

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string)  {
	time.Sleep(1 * time.Second)
	channel <- "arifin"
}

func OnlyOut(channel <-chan string, data *string)  {
	*data = <- channel
}

func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string,2)
	defer close(channel)

	channel <- "Nur"
	channel <- "Arifin"

	data1 := <- channel
	data2 := <- channel
	
	fmt.Println(data1)
	fmt.Println(data2)
}

func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke "+strconv.Itoa(i)
		}
		close(channel)
		fmt.Println("========close==========")
	}()

	for v := range channel {
		fmt.Println("menerima data " + v)
	}
	fmt.Println("===============================SELESAI==========================")

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nur Arifin"
}

func TestSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	var channelNumber int
	for {
		select{
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			channelNumber++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			channelNumber++
		default:
			fmt.Println("Menunggu Data")
		}
		if channelNumber == 2{
			break;
		}
	}
}