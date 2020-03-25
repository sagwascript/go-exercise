package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	fmt.Println(inc.Increment())
	fmt.Println(inc.Increment())
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
