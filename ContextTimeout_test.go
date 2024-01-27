package gopzn

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func createCounter2(ctx context.Context) (chan int, error) {
	destination := make(chan int)
	var err error
	go func () error {
		counter := 1
		for {
			fmt.Println("aku", counter)
			select{
			case <-ctx.Done():
				err = ctx.Err()
				return ctx.Err()
			default:
				destination <- counter
				counter++
				time.Sleep(1*time.Second)
			}
		}
	}()


	return destination,err
}

func TestContextTimeout(t *testing.T) {
	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Millisecond)
	defer cancel()
	
	
	destination,err := createCounter2(ctx)

	
	fmt.Println(<- destination)
	fmt.Println(err)
	for v := range destination {
		fmt.Println("counter ", v)
	}
	
	
	time.Sleep(1*time.Nanosecond)

	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
}

func f(ctx context.Context) {    
	deadline,_:=ctx.Deadline()
	fmt.Println(time.Until(deadline))
}
