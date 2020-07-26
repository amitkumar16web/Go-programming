package main

import "fmt"

func main() {
	strElmntArr := []string{"India", "Assam", "Hydrabad", "Delhi"}
	for i, e := range strElmntArr {
		fmt.Printf("Index Values: %d Element Values: %s\n", i, e)
	}

	m1 := map[string]int{"Amit": 23, "Hafiz": 22, "Kushal": 24}
	for k, v := range m1 {
		fmt.Printf("Key: %s Values: %d\n", k, v)
	}

	for key := range m1 {
		fmt.Println(key)
	}

}
