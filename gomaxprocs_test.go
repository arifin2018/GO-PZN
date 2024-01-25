package gopzn

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T)  {
	totalCPU := runtime.NumCPU()
	fmt.Println("total CPU", totalCPU)
	
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread)
	
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Thread", totalGoroutine)

}