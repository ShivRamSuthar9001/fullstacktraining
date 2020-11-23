//go run command for run file
package main

import "fmt"

// structs
type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	// three type of structs decleration
	shiv := person{"Shiv", "Suthar", contactinfo{"shiv@g.c", 123456}}
	// govind := person{firstName: "govind", lastName: "Singh", contact: contactinfo{email: "govind@g.c", zipCode: 123456}}
	// var kanwar person
	// kanwar.firstName = "Kanwar"
	// kanwar.lastName = "Singh"
	// kanwar.contact = contactinfo{"kanwar@g.c", 1235}
	// shiv.print()

	// pointer call for value change

	// shivPointer := &shiv
	// shivPointer.updateName("shivram")

	//POINTER shoutcut

	shiv.updateName("shivram")
	shiv.print()
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPersson *person) updateName(newFirstName string) {
	(*pointerToPersson).firstName = newFirstName
}
