// Data structure.  Collection of properties that are related together.
// It's like an object in oriented programming language
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//Tell go what fields the person struct has
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//Create a new value of type person
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alexanderson@gmail.com",
			zipCode: 90589,
		},
	}
	alex.updateNamePointer("Test first name")
	alex.updateLastName("Test last name")

	// In go you not directly update the value
	// if you want to update one field you have to alter the pointer of the field
	// alexPointer := &alex    - This notation get the memory adress of this value

	alex.print()

}

//In this function, it receives a pointer to person and update the value to this pointer
//which is the value stored in this memory adress
// this happen in all value types(int,float,string,bool and structs)
// in reference types, you can set without pointer(slice,maps, channels,pointers and functions)

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// if you change it direct , it will change the copy in other memory address so the value won't be affected

func (p person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
