package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	// "%+v" will print out all the field names and values from alex.
}