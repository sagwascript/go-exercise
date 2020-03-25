package main

import "fmt"

func main() {
	cube := Cube{5}
	var s Shape = &cube // assign pointer
	fmt.Println(s.Area())
	fmt.Println(s.Volume()) // this works though it receives pointer instead of value
}

type Shape interface {
	Area() float64
	Volume() float64
}

type Cube struct {
	side float64
}

func (c *Cube) Area() float64 {
	return c.side * c.side
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
