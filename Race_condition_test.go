package gopzn

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T)  {
	var x int

	for i := 0; i < 1000; i++ {
		go func ()  {
			for i := 0; i < 100; i++ {
				x += 1
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(x)
}