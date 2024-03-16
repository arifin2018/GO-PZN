package main

import "fmt"

func main() {
	var name string
	var alphabet_count int

	// Calling Scan() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scan(&name)
	fmt.Scan(&alphabet_count)

	// Printing the given texts
	fmt.Printf("The word %s containing %d number of alphabets.",
		name, alphabet_count)
}
