package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it zero length.
  s = s[:0] // get an empty slice
  printSlice(s) // cap still 6

  // Extend its length.
  s = s[:6] // still can slice more other items later on
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)

  // s = s[:6] // error, current cap is 4 because two items were dropped
  // printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
