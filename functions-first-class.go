package main

import (
  "fmt"
)

func main() {
  func() { // anonymous function
    msg := "Hello Go!"
    fmt.Println(msg)
  }() // called immediately

  for i := 0; i < 5; i++ {
    func(i int) { // best practice, pass through argument
      fmt.Println(i)
    }(i)
  }

  var f func() = func() { // function declaration
    fmt.Println("Hello Go!")
  }
  f()
}
