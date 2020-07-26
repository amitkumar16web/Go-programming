package main

import "fmt"

func main() {

	// eg =1 start
	func(name string) {
		fmt.Println(name)
	}("Hafiz")
	// eg =1 end
	kumsg := ku()
	kumsg("Welcome to Kaziranga University :)")

	seqNo1 := number()
	//fmt.Println(seqNo)
	fmt.Println(seqNo1())
	fmt.Println(seqNo1())
	fmt.Println(seqNo1())

	seqNo2 := number()
	if seqNo2() >= 1 {
		fmt.Println("Hello")
	} else {
		fmt.Println("No")
	}

}

func ku() func(string) {

	return func(msg string) {
		fmt.Println(msg)
	}

}

func number() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}
