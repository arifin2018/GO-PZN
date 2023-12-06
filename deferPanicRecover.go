package main

import (
	"fmt"
)
func endMain()  {
	fmt.Println("End Main")
	message := recover(); // mengambil nilai panic dan membuat program akan berjalan lagi
	fmt.Println("hai,",message)
}


func main()  {
	defer endMain() //defer untuk memanggil function terakhir,dari pada di taro bawah better di taro atas pake defer
	if true {
		panic("lol")// panic sendiri itu untuk memberhentikan program
	}
	endMain() //kalo tidak pake defer dan error dia tidak akan di panggil
}
