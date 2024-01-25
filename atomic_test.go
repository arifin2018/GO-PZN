package gopzn

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroupAndMutex(t *testing.T)  {
	var x int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Counter = ", x)
}

func TestAtomic(t *testing.T)  {
	var x int64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Counter = ", x)
}