package main

import (
	"fmt"
	"os"
)

func main()  {
	args := os.Args
	for _, v := range args {
		fmt.Printf("%s\n", v)
	}

	// go run args.go nur arifin
}