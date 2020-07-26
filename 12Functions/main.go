package main

import "fmt"

func main() {
	fmt.Println("Start @ main...")

	hello()

	e := add(5, 20)
	fmt.Println("a+b=>", e)

	fmt.Println("Ends @ main...")
}

//Create User Defined Function
func hello() {
	fmt.Println("Hello Friends :)")
}

//Functions With Arguments & Return Values
func add(a, b int) int {
	c := a + b
	d := a - b
	fmt.Println("a-b=>", d)
	return c
}
