package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)

	arg1 := os.Args[1]

	fmt.Println(arg1)
}
