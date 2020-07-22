package main

import (
	"fmt"
)

func main() {

	fmt.Print("Please enter a number:")
	var day int
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Print("Moday")
		break
	case 2:
		fmt.Print("Tuday")
		break
	case 3:
		fmt.Print("Wedday")
		break
	case 4:
		fmt.Print("Thday")
		break
	case 5:
		fmt.Print("Frday")
		break
	case 6:
		fmt.Print("saday")
		break
	case 7:
		fmt.Print("Suday")
		break
	default:
		fmt.Print("Nothing")
	}

}
