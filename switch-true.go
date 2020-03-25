package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  // switch without condition is like switch true
  switch {
    // true false condition
    case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
      fmt.Println("Good evening.")
  }

  fmt.Println(t.Hour())
}
