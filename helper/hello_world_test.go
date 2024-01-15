package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("arifin")

	if result != "Hello arifin" {
		// errors.New("INI ERROR")
		panic("error")
	}
}