package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Error nilai pembagi")
	} else {
		return nilai/pembagi , nil
	}
}

func main() {
	// pembagi1,err := pembagian(32,0)
	pembagi1,err := pembagian(32,2)
	if err != nil {
		fmt.Println(err) //Error nilai pembagi
	}else {
		fmt.Println(pembagi1) //16
	}
}