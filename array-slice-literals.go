package main

import "fmt"

func main() {
  // declaring slice of array, like an array without length
  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  // slice array of struct
  s := []struct{
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)

  // slice default just like in Python
  nums := []int{2, 3, 5, 7, 11, 13}
  // all has the same results
  fmt.Println(nums[0:6])
  fmt.Println(nums[:6])
  fmt.Println(nums[0:])
  fmt.Println(nums[:])
}
