package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// rohan := person{"Rohan", "Sundar"}
	// vivek := person{firstName: "Vivek", lastName: "Sundar"}

	// fmt.Println(rohan)
	// fmt.Println(vivek)

	// var rohan person

	// rohan.firstName = "Rohan"
	// rohan.lastName = "Sundar"
	// fmt.Printf("%+v", rohan)
	// fmt.Println("")

	rohan := person{
		firstName: "Rohan",
		lastName:  "Sundar",
		contact: contactInfo{
			email:   "rohan@hotmail.com",
			zipCode: 10000,
		},
	}
	rohan.updateName("Vivek")
	rohan.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
