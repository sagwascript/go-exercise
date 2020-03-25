package main

import "fmt"

func main() {
	cube := Cube{5}
	var s Shape = cube
	var o Object = cube

	// fmt.Println(s.side) // won't work since Shape has no property side
	fmt.Println(s.Area())
	// fmt.Println(s.Volume()) // won't work Shape has no method Volume
	fmt.Println(o.Volume())
	// fmt.Println(o.Area()) // won't work Object has no method Area

	// but Cube has both methods
	fmt.Println(cube.Area())
	fmt.Println(cube.Volume())

	// use type assertion to extract dynamic value of an interface
	extractedCube := s.(Cube) // it's a Cube, NOT an interface which has dynamic value of Cube
	fmt.Println("extractedCube area:", extractedCube.Area())
	fmt.Println("extractedCube volume:", extractedCube.Volume()) // now it has Volume method
	fmt.Println("extractedCube side: ", extractedCube.side)      // it's available too

	// type asserting non implemented interface
	if skin, ok := s.(Skin); ok {
		fmt.Println(skin)
	} else {
		fmt.Println("cube isn't implement Skin")
	}
}

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Skin interface {
	Color() float64
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
