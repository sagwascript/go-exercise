package main

import "fmt"

// type annotation
func add(x, y int) int {
  return x + y
}

// return multiple value
func swap(x, y string) (string, string) {
  return y, x
}

// naked return to return named return values
func split(num int) (x, y int) {
  x = num * 4 / 9
  y = num - x
  return
}

func main() {
  fmt.Println(add(1, 5))
  a, b := swap("hello", "world")
  fmt.Println(a, b)
  fmt.Println(split(17))
}
