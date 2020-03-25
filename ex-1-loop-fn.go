package main

import (
	"fmt"
    "math"
)

func Sqrt(x float64) float64 {
  z := x/2
  stop := false
  for i := 1; !stop; i++ {
    z -= (z*z - x) / (2*z)
    if z*z - x  < 0.0000000001 {
      stop = true
    }
  }
  return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(5))
	fmt.Println(Sqrt(25))
    fmt.Println("--")
    fmt.Println(math.Sqrt(2))
    fmt.Println(math.Sqrt(3))
    fmt.Println(math.Sqrt(4))
    fmt.Println(math.Sqrt(5))
}

