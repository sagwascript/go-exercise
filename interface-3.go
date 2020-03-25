package main

import "fmt"

func main() {
	cube := Cube{5}
	var m Material = cube
	var s Shape = cube
	var o Object = cube
	fmt.Printf("dynamic type and value of interface m of static type Material is %T and %v\n", m, m)
	fmt.Printf("dynamic type and value of interface s of static type Shape is %T and %v\n", s, s)
	fmt.Printf("dynamic type and value of interface o of static type Object is %T and %v\n", o, o)
	fmt.Println("Area:", m.Area()) // interface m provide both methods
	fmt.Println("Volume:", m.Volume())
}

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

// embed/compose other interface
type Material interface {
	Shape
	Object
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return c.side * c.side
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
