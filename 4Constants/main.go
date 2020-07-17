package main

import "fmt"

func main() {
	//Go supports constants of character, string, boolean, and numeric values.
	const a string = "Amit" // not change
	const b = 10
	// b = 12 error

	fmt.Println(a)
	fmt.Println(b)
}
