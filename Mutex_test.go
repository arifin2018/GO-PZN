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

	// time.Sleep(2 * time.Second)
	fmt.Println("Total balance", data)
}

type UserBalance struct {
	sync.Mutex
	Name	string
	Balance	int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int)  {
	user1.Mutex.Lock()
	fmt.Println("Lock User 1", user1.Name, " - See Amount", amount)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Mutex.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T)  {
	user1 := UserBalance{
		Name: "Nur",
		Balance: 1000,
	}
	user2 := UserBalance{
		Name: "Arifin",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(1 * time.Second)

	fmt.Println("User ", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, " Balance ", user2.Balance)
}

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