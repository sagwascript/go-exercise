package main

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
  // variable v will not available outside the if scope
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim 
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(4, 4, 1000),
  )
}
