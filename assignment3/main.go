package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type consoleWriter struct{}

func main() {

	// fmt.Println(os.Args)
	arg1 := os.Args[1]

	data, err := ioutil.ReadFile(arg1)

	cw := consoleWriter{}
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(cw, data)
}

func (consoleWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Printed out number of bytes: ", len(bs))
	return len(bs), error
}
