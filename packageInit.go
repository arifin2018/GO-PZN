package main

import (
	"GO-PZN/database"
	_ "GO-PZN/internal" // underscore == blank identifier
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
	
	// output
	// this is internal
	// pgsql
}