package main

import "fmt"

type address struct {
	name string
}


func main()  {
	address1 := address{"jakarta"}
	address2 := &address1

	address2.name = "arifin"//kalo seperi ini address1 dan address2 akan berubah menjadi arifin

	address2 = &address{"jogja"}//kenapa menunggunakan dan? karna sebelumnya address2 adalah pointer dari address 1
	
	fmt.Println(address1) //arifin
	fmt.Println(address2) //arifin
	
	address2.name = "bandung"
	fmt.Println(*address2)//bandung

	//asterisk operator
	address2 = &address1 //kenapa perlu di panggil lagi?karena address2 sudah sempat meniban pointer dengan &address{"jogja"}
	*address2 = address{"sulawesi"}
	fmt.Println(address1)//sulawesi
	fmt.Println(address2)//sulawesi

}