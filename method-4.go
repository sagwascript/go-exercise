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
	john.contact.changePhoneNumber("0812 4432 5323") // call method on struct field
	fmt.Println(john)
}

type contact struct {
	phone, address string
}

type employee struct {
	name    string
	salary  int
	contact contact
}

func (c *contact) changePhoneNumber(newNumber string) {
	c.phone = newNumber
}
