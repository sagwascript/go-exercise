package main

import "fmt"
import "runtime"

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("macOS")
  case "linux":
    fmt.Println("Linux")
  default:
    fmt.Printf("%s.\n", os)
  }

  switch i := 2 + 3;i { // called initializer
  case 1, 5, 10: // multiple test because switch is NOT implicitly falling through in Go
    fmt.Println("one, five or ten!")
  case 2, 4, 6:
    fmt.Println("two, four or six")
  default:
    fmt.Println("another number")
  }

  j := 10
  switch { // tag-less switch
  case j <= 10:
    fmt.Println("less than or equal 10")
    fallthrough // to make it fallthrough, beware it doesn't care the condition next case
  case j <= 20:
    fmt.Println("less than or equal 20")
  default:
    fmt.Println("greater than 20")
  }
}
