package main

import "fmt"

func main() {
  var s []int // create empty slice
  fmt.Println(s, len(s), cap(s))
  // return true
  if s == nil {
    fmt.Println("nil!")
  }
}
