package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// fmt.Println(os.Args)
	arg1 := os.Args[1]

	file, err0 := os.Open(arg1)
	if err0 != nil {
		fmt.Println("Error: ", err0)
		os.Exit(1)
	}

	bs := make([]byte, 9999)
	data, err := file.Read(bs)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sr := strings.NewReader(string(bs[:data]))
	io.Copy(os.Stdout, sr)
}
