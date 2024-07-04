package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.contactInfo = contactInfo{email: "hello@gmail.com", zip: 95118}
	alex.updateName("asdasd")

	alex.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}
