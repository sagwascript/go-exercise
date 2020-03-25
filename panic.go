package main

import (
  "fmt"
)

func main() {
  // a, b := 1, 0
  // ans := a/b // this will cause panic
  // fmt.Println(ans)

  // commented number is order of execution
  fmt.Println("start") // 0
  defer fmt.Println("this was deferred") // 1
  panic("something wrong") // 2
  fmt.Println("end") // 0
}
