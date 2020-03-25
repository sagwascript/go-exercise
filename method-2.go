package main

import "fmt"

func main() {
	john := employee{name: "John Doe", salary: 1000}
	fmt.Println("Before changed", john)
	john.changeName("Mark")
	fmt.Println("After changed", john) // John is still John because receiver copy the value before modifies it

	matthew := employee{name: "Matthew Doe", salary: 500}
	fmt.Println("Before changed", matthew)
	(&matthew).changeNamePointer("Mark Doe") // use pointer as a receiver of the method
	fmt.Println("After changed", matthew)    // yay, it changes!

	luke := employee{name: "Luke Doe", salary: 750}
	fmt.Println("Before changed", luke)
	luke.changeNamePointer("Jane Doe") // no need to use addressof operator on luke, Go figures that out
	fmt.Println("After changed", luke) // it still changes
}

type employee struct {
	name   string
	salary int
}

func (e employee) changeName(newName string) {
	e.name = newName // change receiver property
}

// accept pointer receiver
func (e *employee) changeNamePointer(newName string) {
	e.name = newName
}
