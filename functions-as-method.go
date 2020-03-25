package main

import (
  "fmt"
)

func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }
  g.greet()
}

type greeter struct {
  greeting string
  name string
}

// method
func (g greeter) greet() { // that struct copied, except you use pointer
  fmt.Println(g.greeting, g.name)
  g.name = "" // this has no effect
}
