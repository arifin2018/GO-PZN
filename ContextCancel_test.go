package gopzn

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextCancel(t *testing.T) {
	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destinationChan := make(chan int)
	destination := createCounter(destinationChan,ctx)
	
	// fmt.Println(<- destination)
	for v := range destination {
		fmt.Println("counter ", v)
		if v == 10 {
			break;
		}
	}
	cancel()
	time.Sleep(1*time.Nanosecond)


	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
}

func createCounter(destination chan int, ctx context.Context) chan int {
	go func ()  {
		counter := 1
		for {
			fmt.Println("aku", counter)
			select{
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++	
			}
		}
	}()


	return destination
}

func TestChannelMake(t *testing.T) {
	channel := Channel()
	for v := range channel {
		if v == 10 {
			break;
		}
	}
}

func Channel() chan int {
	channel := make(chan int)
	go func ()  {
		for i := 0; i < 23; i++ {
			channel <- i
			fmt.Println("Awd", i)
		}
	}()
	return channel
}