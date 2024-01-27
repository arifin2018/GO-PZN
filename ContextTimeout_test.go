package gopzn

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()
	destinationChan := make(chan int)
	destination := createCounter(destinationChan,ctx)
	
	fmt.Println(<- destination)
	for v := range destination {
		fmt.Println("counter ", v)
		if v == 2 {
			break;
		}
	}
	time.Sleep(1*time.Nanosecond)

	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
}