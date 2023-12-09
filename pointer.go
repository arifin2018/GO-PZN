package main

import "fmt"

type address struct {
	name string
}

func main()  {
	address1 := address{name: "jakarta"}
	address1_1 := address1
	address1_1.name = "Jogja"

	address2 := address{name: "bandung"}
	address2_1 := &address2
	address2_1.name = "Surakarta"

	fmt.Println(address1) //{jakarta} 
	fmt.Println(address1_1) //{Jogja} (disini kita bisa melihat bahwa data address1nya masih jogja belum berubah menjadi jogja,karena data yang di address1_1 hanya mengcopy data di address1)

	fmt.Println(address2) //{Surakarta} 
	fmt.Println(address2_1) //&{Surakarta} (disini kita bisa lihat bahwa ketika kita menggunakan dan data di address2 nya juga berubah karena datanya pass by reference yang artinya data memory nya di ambil di taruh di address2_1 juga,jadi ketika address2_1 berubah automatis data di address2 juga ikut berubah)
	fmt.Println(*address2_1) //{Surakarta}
}