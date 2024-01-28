package gopzn

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
Sumber : https://blog.dot.co.id/mengenal-context-pada-bahasa-pemrograman-go-f198b2f4b71f
*/

func calculate(ctx context.Context, x int, y int) int {
	// melakukan perhitungan matematika sederhana
	sum := x + y
	time.Sleep(2*time.Second)

	// menunggu selama 10 detik untuk mensimulasikan suatu proses yang membutuhkan waktu lama
	select {
	case <-ctx.Done():
	 // jika context dibatalkan, kembalikan nilai -1
		return -1
	default:
	 // jika tidak, kembalikan hasil perhitungan
		return sum
	}
}

func TestDigitalOcean(t *testing.T)  {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// menjalankan fungsi yang membutuhkan context
	result := calculate(ctx, 10, 5)
	fmt.Println(result)
}