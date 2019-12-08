package main

import "fmt"

func main() {

	sliceOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range sliceOfInts {
		if value%2 != 0 {
			fmt.Println(value, " is odd")
		} else {
			fmt.Println(value, " is even")
		}
	}
}
