package main

import "fmt"

func main() {
	var i interface{} = [...]int{1, 2, 3}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}
