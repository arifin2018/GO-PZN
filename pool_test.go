package gopzn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T)  {
	pool := sync.Pool{
		New : func() any {
			return "name"
		},
	}

	pool.Put("Nur")
	pool.Put("Arifin")

	for i := 0; i < 3; i++ {
		go func ()  {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}

type Person struct {
    Name string
}

var pool = sync.Pool{
    New: func() any {
        fmt.Println("Creating a new person...")
        return &Person{
			Name: "arifin",
		}
    },
}

func TestMainPool(t *testing.T) {
    person := pool.Get().(*Person)
    fmt.Println("Get object from sync.Pool for the first time:", person)

    fmt.Println("Put the object back in the pool")
    pool.Put(person)

    person.Name = "Gopher"
    fmt.Println("Set object property name:", person.Name)

    fmt.Println("Get object from pool again (it's updated):", pool.Get().(*Person))
    fmt.Println("There is no object in the pool now (new one will be created):", pool.Get().(*Person))
}