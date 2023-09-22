package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo // We can omit var name if it is the same as the type name.
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim.party@gmail.com",
			zipCode: 94000,
		}, // All lines need commas
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}