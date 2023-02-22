package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // can also just be contactInfo (w/o var name like contact)
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{firstName: "Jim", lastName: "Carry", contact: contactInfo{
		email:   "jim@gmail.com",
		zipCode: 94000,
	},
	}

	jim.updateName("jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
