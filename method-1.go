package main

import "fmt"
import "math"

func main() {
	john := person{firstname: "John", lastname: "Doe"}
	fmt.Println(john.getFullName())
	rectangle := rect{width: 12, height: 5}
	circ := circle{radius: 7}
	fmt.Println("Area of rectangle is", rectangle.area())
	fmt.Println("Area of circle is", circ.area())
}

type person struct {
	firstname, lastname string
}

func (p person) getFullName() string {
	return p.firstname + " " + p.lastname
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

// you can have method with the same name as long as it has different receiver type
func (c circle) area() float64 {
	return math.Pi * c.radius
}
