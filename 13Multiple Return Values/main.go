package main

import "fmt"

func main() {
	e, f := addMul(5, 10)
	fmt.Println("a+b=>", e)
	fmt.Println("a*b=>", f)

	name, age := nameAge("Amit", 22)
	fmt.Println("Name= ", name)
	fmt.Println("Age= ", age)

	_, id := notUse()
	fmt.Println("id=> ", id)

}

//multiple return values

func addMul(a, b int) (int, int) {
	c := a + b
	d := a * b
	return c, d
}

func nameAge(n string, z int) (string, int) {
	m := n
	g := z
	return m, g

}

func notUse() (string, int) {
	return "Hafiz", 63

}
