package main

import "fmt"

func main() {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  
  var s []int = primes[1:4] // slice from index 1 through 3
  fmt.Println(s)
  // s[3] = 17 // not work, slice has length of 3

  s[1] = 27 // slice reference original array
  fmt.Println(primes) // primes change
}
