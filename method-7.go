package main

import "fmt"

func main() {
	john := employee{
		name:   "John Doe",
		salary: 1000,
	}
	fmt.Println(john)
	john.changeName("Mark Doe") // works though we call it using type value (not pointer)
	john.showSalary()           // printed salary changed but salary of struct john doesn't because it's copied
	(&john).showSalary()        // try pass pointer instead of value of john, still won't change because the method expect a value rather than pointer
	fmt.Println(john)
}

type employee struct {
	name   string
	salary int
}

func (e *employee) changeName(newName string) {
	e.name = newName // phone can be accessed because it's promoted in employee struct
}

func (e employee) showSalary() {
	e.salary = 1500 // bad practice, just for example
	fmt.Println("Salary is", e.salary)
}
