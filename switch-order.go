package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  // it stops when the case is true and skip the rest
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }
}
