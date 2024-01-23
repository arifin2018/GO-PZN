package gopzn

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld(name string)  {
	time.Sleep(1 * time.Second)
	fmt.Println("hello ", name)
}

func TestHelloWorld(t *testing.T)  {
	go HelloWorld("arifin")
	fmt.Println("arifin")
	
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int)  {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T)  {
	for i := 0; i < 100; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(1 * time.Second)
}