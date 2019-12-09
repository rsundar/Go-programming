package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	rohan := person{"Rohan", "Sundar"}
	vivek := person{firstName: "Vivek", lastName: "Sundar"}

	fmt.Println(rohan)
	fmt.Println(vivek)
}
