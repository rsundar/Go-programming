package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// rohan := person{"Rohan", "Sundar"}
	// vivek := person{firstName: "Vivek", lastName: "Sundar"}

	// fmt.Println(rohan)
	// fmt.Println(vivek)

	var rohan person

	rohan.firstName = "Rohan"
	rohan.lastName = "Sundar"
	fmt.Printf("%+v", rohan)
	fmt.Println("")
}
