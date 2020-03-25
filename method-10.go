package main

import "fmt"

func main() {
	number := 1
	var yNum yoloNum = yoloNum(number)
	// var yNum yoloNum = 1 // or you can simply do this

	fmt.Println(yNum)   // output: 1
	(&yNum).increment() // it works
	yNum.increment()    // it works as well
	yNum.increment()
	fmt.Println(yNum) // output: 4
}

type yoloNum int

func (n *yoloNum) increment() {
	*n++
}
