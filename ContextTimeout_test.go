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
	ctx, cancel := context.WithTimeout(parent, 3*time.Nanosecond)
	defer cancel()
	destinationChan := make(chan int)
	
	destination := createCounter(destinationChan,ctx)
	
	fmt.Println(<- destination)
	for v := range destination {
		select {
		case <-time.After(1*time.Nanosecond):
			break
		case <-ctx.Done():
			break
		}
		fmt.Println("counter ", v)
		if v == 29 {
			break;
		}
	}
	
	
	time.Sleep(1*time.Nanosecond)

	fmt.Println("Total Golang Goroutine", runtime.NumGoroutine())
}

func f(ctx context.Context) {    
	deadline,_:=ctx.Deadline()
	fmt.Println(time.Until(deadline))
}
