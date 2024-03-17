package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, data chan<- string) {
	select {
	case <-time.After(1 * time.Second): // Simulasi operasi yang memakan waktu
		data <- "Data berhasil diambil"
	case <-ctx.Done():
		data <- ctx.Err().Error() + " data error fetchData"
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dataChan := make(chan string)
	go fetchData(ctx, dataChan)

	select {
	case result := <-dataChan:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Gagal mengambil data:", ctx.Err())
	}
}
