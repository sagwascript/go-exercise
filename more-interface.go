package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// you can assign geometry type variable to rectangle since it implements the interface
	var shape geometry = rectangle{height: 20, width: 5}
	rectangle := rectangle{height: 20, width: 5}
	fmt.Println(shape)
	fmt.Printf("Type of shape is %T\n", shape)
	fmt.Printf("Value of shape is %v\n", shape)
	fmt.Println("Area of rectangle is", shape.area())
	fmt.Println("shape == rectangle is", shape == rectangle)
	// reassign to other type that implements geometry, circle
	shape = circle{radius: 4}
	fmt.Println(shape)
	fmt.Printf("%T %v\n", shape, shape)

	explain(shape)
	explain(rectangle)
	explain("something")
	explain(19)
}

// this function can accept any kind of type, since any kind of type implements empty interface
func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string))) // type assertion on i because method ToUpper only accept type string
	case int:
		fmt.Println("i stored int", i) // no need to type assert since Println accept anything
	default:
		fmt.Println("i stored something else", i)
	}
}

type geometry interface {
	area() float64
	perim() float64
}

// implicitly implement the interface, you should create exactly those two methods in interface
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}
