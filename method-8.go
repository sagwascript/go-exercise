package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello world"
	fmt.Println(hello.toUpperCase()) // supposed to work, but it's not
}

// this won't work because type string and this method is located on different package
func (s string) toUpperCase() string {
	return strings.ToUpper(s)
}
