package main

import "fmt"

func main() {
	mvalues := mul(1, 2, 3, 4)
	fmt.Println("Return tatal values=> ", mvalues)

	//slices
	sls := []string{"Hafiz", "Hunnan", "Najmin", "Namrat"}
	allname(sls...)

}

func mul(nums ...int) int {
	total := 1

	for _, n1 := range nums {
		fmt.Println("n1=>", n1)
		total *= n1
		fmt.Println("Out side of main func> ", total)

	}
	return total

}

func allname(name ...string) {

	for i, v := range name {
		fmt.Println("Index= ", i)
		fmt.Println("Values= ", v)

	}

}
