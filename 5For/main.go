package main

import "fmt"

func main() {
	for b := 1; b <= 5; b++ {
		fmt.Println(" ", b)

	}
	for c := 1; c <= 5; c++ {
		if c%2 == 0 {
			continue
		}
		fmt.Println(" ", c)
	}

}
