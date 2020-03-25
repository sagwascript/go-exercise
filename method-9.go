package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello yoloString = "Hello world"
	fmt.Println(hello.toUpperCase()) // it works now
}

// create derived type of string to make method toUpperCase works
type yoloString string

// this will work
func (s yoloString) toUpperCase() string {
	return strings.ToUpper(string(s)) // cast s to string, since ToUpper is only accept type string NOT yoloString
}
