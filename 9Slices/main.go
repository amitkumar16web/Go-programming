package main

import "fmt"

func main() {
	a := make([]string, 3)
	fmt.Println(len(a))
	fmt.Println(a)

	a[0] = "amit"
	a[1] = "kumar"
	a[2] = "singh"
	fmt.Println(a)
	fmt.Println(len(a))
	a = append(a, "A")
	a = append(a, "B")
	a = append(a, "C")
	fmt.Println("New elmnt:>", a)
	fmt.Println(len(a))
	c := make([]string, len(a))
	copy(c, a)
	fmt.Println("Copy elmnt:>", c)
	fmt.Println(len(c))
	l := a[2:5]
	fmt.Println("slice", l)

}
