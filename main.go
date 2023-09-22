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

	jimPointer := &jim
	jimPointer.updateName("Jimmy")

	// pointer shortcut
	// if receiver function has *person, jim is automatically received as &jim.
	// jim.updateName("Jimmy")

	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// p is a copy of "jim" that points to a different memory location
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// &jim - memory address (pointer)
// provides the memory address where jim is pointing at.

// *pointerToPerson - value
// provides the value to which this memory adress is pointing at.

// *person - type: pointer to a person
// different from above two! Declares that the type is a "pointer to a person"

// Note: This only applies to structs.
// No need to worry for slices.
// This is because slices are reference data structures that point to an underlying array.
// The slices are copied, but the underlying array is the same, and both copies point to the same array.