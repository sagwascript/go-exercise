package main

import "fmt"

func main() {
  v := 42i
  // reassigning with other type will cause error
  // v = "a"
  fmt.Printf("v is type of %T\n", v)
}
