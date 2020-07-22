package main

import "fmt"

func main() {

	var a [6]int //   0 1 2 3 4 5 = total 6 elmnts  n-1     6-1=5
	fmt.Println("elmnt", a)

	//a[6] = 20 // error
	a[5] = 50 // n-1 6-1= 5
	fmt.Println("elmnt", a)
	fmt.Println("len", len(a))

	b := [3]int{10, 20, 30}
	fmt.Println("All emnt of b", b)

	var tdarray [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tdarray[i][j] = i + j
		}
	}
	fmt.Println("2D elmnt=>", tdarray)

}
