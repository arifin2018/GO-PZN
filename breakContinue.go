package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue; // jika angka nya habis di bagi 2,lewati saja,atau continue saja
		}
		fmt.Println(i)
	}
	fmt.Println("============================================================================")
	for i := 0; i < 10; i++ {
		if i == 7 {
			break; // jika i nya sama dengan 7,maka berhentikan saja di dalam kode block loop ini
		}
		fmt.Println(i)
	}
}