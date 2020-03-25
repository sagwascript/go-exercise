package main

import "fmt"

func main() {
  defer fmt.Println("world")
  fmt.Println("hello")

  // defer typically used for closing opened resource
  // defer is LIFO, defer is like pushing something to event loop in js
  // it will get executed while this function done executing non-defer statements
  fmt.Println("Stacking defer")
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done.")

  a := "start"
  // arguments evaluated at time defer is executed, not when the function executed
  defer fmt.Println(a) // a will evaluated to "start"
  a = "end"
}
