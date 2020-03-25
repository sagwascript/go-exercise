package main

import "fmt"
import "math"

func sqrt(x float64) string {
  // no need to use parentheses (optional)
  if x < 0 {
    return sqrt(-x) + "i" // as it's a complex number
  }
  return fmt.Sprint(math.Sqrt(x))
}

func main() {
  fmt.Println(sqrt(2), sqrt(-9))
  statePopulations := map[string]int{
    "CA": 6234342,
    "TX": 6234343,
    "FL": 8323444,
    "NY": 3324252,
    "PN": 4234235,
    "IL": 2342523,
    "OH": 5234235,
  }
  if pop, ok := statePopulations["NY"]; ok {
    fmt.Println(pop) // pop can only accessed in this scope
  }
}
