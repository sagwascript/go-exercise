package main

import (
  "fmt"
  "log"
)

func main() {
  fmt.Println("start")
  panicker()
  fmt.Println("end") // will still execute this because of recover
}

// use recover in deferred function
func panicker() {
  fmt.Println("about to panic")
  defer func() {
    if err := recover(); err != nil {
      log.Println("Error: ", err)
      // panic(err) // when the program can't deal with that error than re-panic to stop any execution
    }
  }()
  panic("trigger panic")
  fmt.Println("done panicking") // it won't execute this
}
