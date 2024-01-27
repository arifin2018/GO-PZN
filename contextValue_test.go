package gopzn

import (
	"context"
	"fmt"
	"testing"
)

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")


	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)


	fmt.Println("=========================================CONTEXT GET VALUE===================================")
	fmt.Println(contextF.Value("f")) //berhasil ambil data context karena masih 1 turunan atau sama datanya
	fmt.Println(contextF.Value("c")) //berhasil ambil data context karena masih 1 turunan dari parent nya atau sama datanya
	fmt.Println(contextF.Value("b")) //tidak berhasil ambil data context karena sudah beda turunan atau beda datanya
	fmt.Println(contextA.Value("b")) //tidak berhasil ambil data context karena parent tidak bisa ambil data child
}