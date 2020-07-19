package main

import "fmt"

func main() {
	/*
		fmt.Print("Please enter a number >")
		var n int
		fmt.Scan(&n)
		if n%2 == 0 {
			fmt.Println(n, " is an even number")
		} else {
			fmt.Println(n, " is an old number")
		}
	*/
	/*
		fmt.Print("Please enter your 10th % >")
		var m1 int
		fmt.Scan(&m1)
		if m1 >= 60 {
			fmt.Print("Please enter your 12th % >")
			var m2 int
			fmt.Scan(&m2)
			if m2 >= 60 {
				fmt.Print("Yes")
			} else {
				fmt.Print("No")
			}

		} else {
			fmt.Print("No")
		}
	*/
	fmt.Print("Please enter  a number>")
	var num int
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num <= 9 {
		fmt.Println(num, "has 1 digit")
	} else if num < 99 {
		fmt.Println(num, "has 2 digits")
	} else if num == 99 {
		fmt.Println(num, "has 2 digit highest  numbe ")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
