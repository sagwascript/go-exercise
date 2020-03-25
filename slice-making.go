package main

import "fmt"

func main() {
  a := make([]int, 5) // make slice with length of 5
  printSlice(a)

  b := make([]int, 3, 5) // you can't make cap less than the length
  printSlice(b)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
