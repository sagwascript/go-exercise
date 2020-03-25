package main

import "fmt"
import "math"

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  // error should explicitly convert the type
  // var z uint = f
  fmt.Println(x, y, z)
}
