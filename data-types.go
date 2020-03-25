package main

import "fmt"

// declare variable
var c, python, java bool

// declare variable with initial value
var j, k int = 1, 4

func main() {
  var i int

  // omit the type
  var html, css, js = true, false, "yay"

  // short variable declaration in a func (w/o type and only works inside func)
  name := "john"

  fmt.Println(i, c, python, java)
  fmt.Println(j, k)
  fmt.Println(html, css, js)
  fmt.Println(name)
}
