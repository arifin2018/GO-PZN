package main

import "fmt"

func main() {
	/**
	- map adalah tipe data lain yang berisikan kumpulan data yang sama namun kita bisa menentukan key yang bisa kita inginkan
	- tidak ada batasan untuk data map bisa sebanyak banyak nya,asalkan key nya berbeda
	- jika kita menggunakan key yang sama,maka key yang sudah ada datanya akan ketiban
	*/

	person := map[string]string{
		"name":"arifin",
		"address":"jagakarsa",
	}

	fmt.Printf("%s\n", person)//map[address:jagakarsa name:arifin]

	/**
	ini adalah sama cara membuat map kosong

	book := make(map[string]string)
	var book map[string]string
	*/
	book := make(map[string]string)
	book["title"] = "buku arifin"
	book["author"] = "arifin"
	fmt.Println(book) //map[author:arifin title:buku arifin]
	delete(book,"author")
	fmt.Println(book) //map[title:buku arifin]
}