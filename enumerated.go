package main

import "fmt"

// enum is only enumerated in constant block scope, start from 0
const (
  a = iota << 3 // you can use bit shifting to make this dynamic, remember you can't call fn in constant
  b
  c
)

const (
  john = iota
  mark
)

const (
  isAdmin = 1 << iota
  isHeadquarters
  canSeeFinancial
)

func main() {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  // iota reset to 0
  fmt.Println(john)
  fmt.Println(mark)
  fmt.Println("--")
  fmt.Printf("%v\n", isAdmin)
  fmt.Printf("%v\n", isHeadquarters)
  fmt.Printf("%v\n", canSeeFinancial)
}
