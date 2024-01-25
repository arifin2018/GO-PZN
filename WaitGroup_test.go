package gopzn

import (
	"fmt"
	"sync"
	"testing"
)


func RunAsynchronous(group *sync.WaitGroup, number int)  {
	defer group.Done()
	// group.Add(1)
	fmt.Println("Hello ", number)
}

func TestWaitGroup(t *testing.T)  {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RunAsynchronous(&wg, i)
	}

	wg.Wait()
	fmt.Println("Selesai")
}