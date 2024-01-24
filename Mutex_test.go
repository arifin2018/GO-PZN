package gopzn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T)  {
	var x int
	var a int
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func ()  {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
				a += 1
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(x)
	fmt.Println(a)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int)  {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance(data *int) {
	account.RWMutex.RLock()
	*data = account.Balance
	account.RWMutex.RUnlock()
}

func TestRWMutex(t *testing.T)  {
	account := BankAccount{}
	var data int

	for i := 0; i < 10; i++ {
		go func ()  {
			for i := 0; i < 50; i++ {
				account.AddBalance(1)
				account.GetBalance(&data)
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total balance", data)
}