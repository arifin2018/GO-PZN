package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name,os.O_CREATE|os.O_RDWR,0755)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(message))
	return nil
}

func appendNewLineFile(name string,message string) error {
	file, err := os.OpenFile(name,os.O_CREATE|os.O_APPEND,0755)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string,error) {
	file,err := os.OpenFile(name, os.O_RDONLY, 0444)
	if err != nil {
		return "",err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line,_,err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message = message + string(line)
	}
	return message, nil
}

func main() {
	// createNewFile("golang.log","nur arifin")
	// appendNewLineFile("golang.log","\nnur arifin 2")
	read, _ := readFile("golang.log")
	fmt.Println(read)
}