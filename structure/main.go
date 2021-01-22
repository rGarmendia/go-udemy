package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Morrison",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Pointing to the pointer and not the value that is has
	//     Pointer
	//  adress|value
	//  0001  |person{firstName:"juanito",...}

	jimPointer := &jim
	jimPointer.updateFirstName("Juanito")
	jim.print()

}

// To point the function receiver, to a pointer you have to prepend a * to the type,
// it means that you are summoning a POINTER and not its value, example *person
func (pointerToPerson *person) updateFirstName(newFirstName string) {

	// To go back into the value from a pointer perspective, you'd put a pair of () and inside call the instance with an * before it, example (*pointerToPerson)
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
