package main

import "fmt"

func main()  {
	var names = [...]string{"azriel","rafiq","pradipta","khoirunnisa","miftahul","jannah"}
	fmt.Println(names) //[azriel rafiq pradipta khoirunnisa miftahul jannah]
	fmt.Println(names[:3])// pointer 0, capacity 3 ([azriel rafiq pradipta])
	fmt.Println(cap(names[:3]))// length capacity 3 (0,1,2) //(6)
	fmt.Println(names[3:5]) //[khoirunnisa miftahul]
	fmt.Println(names[3:]) //[khoirunnisa miftahul jannah]

	nameSliceKhoirunnisa := names[3:]
	nameSliceKhoirunnisa[1] = "MJ"
	fmt.Println(nameSliceKhoirunnisa) //[khoirunnisa MJ jannah]
	fmt.Println(names) //[azriel rafiq pradipta khoirunnisa MJ jannah]

	data := append(names[:], "Arifin")
	data[0] = "Nur"
	fmt.Println(data)//[azriel rafiq pradipta khoirunnisa MJ jannah Arifin]
	fmt.Println(names)// names pada index 0 tidak berubah karena di append [azriel rafiq pradipta khoirunnisa MJ jannah]

	var newSlice []string = make([]string, 2)// length 2
	//yang artinya length saat ini harus di isi 2 atau default nya harus dulu di isi 2 data,namun sebenernya variable `newSlice` ini masih bisa di tambah sampe 5 akan tetapi nambahnya musti pake append
	newSlice[0] = "nur"
	newSlice[1] = "arifin"
	arrayString := []string{"ganteng","banget","awd","awss","xjxjxj","awd","asu"}
	newSlice2:= append(newSlice, arrayString...)
	fmt.Println(newSlice2)

	iniArray := [...]int{1,2,3,4}
	iniSlice := []int{1,2}
	fmt.Println(iniArray, iniSlice)
}	