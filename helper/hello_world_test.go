package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.M)  {
	fmt.Println("FIRST arifin")
	
	t.Run()
	
	fmt.Println("END arifin")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("can't running in windows,please use another os")
	}
	result := HelloWorld("arifin")
	require.Equal(t, "Hello arifin", result);
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("arifin")

	require.Equal(t, "Hello arifin", result);
}
func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("arifin")

	if result != "Hello arifin" {
		// errors.New("INI ERROR")
		// panic("error")
		t.Error("hai")
	}
}

func TestHelloWorldAssertion(t *testing.T)  {
	result := HelloWorld("arifin")

	assert.Equal(t, "Hello arifin", result)
}