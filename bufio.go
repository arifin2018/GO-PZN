package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/**
	bufio read
	*/
	input := strings.NewReader("name:nur arifin\nage:24 years old")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	/**
		bufio write
	*/
	write := bufio.NewWriter(os.Stdout)
	_,_ = write.WriteString("hello world\n")
	_,_ = write.WriteString("selamat belajar\n")
	write.Flush()
}