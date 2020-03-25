package main

import "fmt"

func main() {
	john := employee{
		name:   "John Doe",
		salary: 1000,
		contact: contact{
			phone:   "0811 1234 4444",
			address: "Indonesia",
		},
	}
	fmt.Println(john)
	fmt.Println("John phone number is", john.phone) // phone is directly accessible
	fmt.Println("John address is", john.address)
	john.changePhoneNumber("0812 4432 5323")
	fmt.Println(john)
}

type contact struct {
	phone, address string
}

type employee struct {
	name    string
	salary  int
	contact // anonymously embed struct, this will make embedded struct merge it's properties
}

func (e *employee) changePhoneNumber(newNumber string) {
	e.phone = newNumber // phone can be accessed because it's promoted in employee struct
}
