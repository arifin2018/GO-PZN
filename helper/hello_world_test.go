package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T)  {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name:"nur",
			request:"nur",
			expected:"Hello nur",
		},
		{
			name:"arifin",
			request:"arifin",
			expected:"Hello arifin",
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := HelloWorld(v.request)
			require.Equal(t, v.expected, result);
		})
	}
}

func TestMain(t *testing.M)  {
	fmt.Println("FIRST arifin")
	
	t.Run()
	
	fmt.Println("END arifin")
}

func TestSubTest(t *testing.T)  {
	t.Run("Nur", func(t *testing.T) {
		result := HelloWorld("arifin")
		require.Equal(t, "Hello arifin", result);
	})
	t.Run("Arifin", func(t *testing.T) {
		result := HelloWorld("arifin")
		require.Equal(t, "Hello arifin", result);
	})
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

func BenchmarkHelloWorld(t *testing.B)  {
	for i := 0; i < t.N; i++ {
		HelloWorld("arifin")
	}
}

func BenchmarkHelloWorldSub(t *testing.B)  {
	t.Run("Nur", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			HelloWorld("nur")
		}
	})
	t.Run("arifin", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			HelloWorld("arifin")
		}
	})
}

func BenchmarkTable(t *testing.B)  {
	Benchmark := []struct{
		name string
	}{
		{
			name: "nur",
		},
		{
			name: "arifin",
		},
	}

	for _, v := range Benchmark {
		t.Run(v.name, func(b *testing.B) {
			for b := 0; b < t.N; b++ {
				HelloWorld(v.name)
			}
		})
	}
}