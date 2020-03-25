package main

import (
"fmt"
"strconv"
)

func test(x int) int {
	return x
}

// uppercase variable will exported and can be accessed by other package
var Job string = "DevOps"

func main() {
	var num int = 62
    var numStr string = strconv.Itoa(num)
	fmt.Println(num)
    fmt.Println(numStr)
    x := 8.1
    fmt.Printf("%T\n", x)
}
