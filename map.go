package main

import (
  "fmt"
)

func main() {
  statePopulations := make(map[string]int)
  statePopulations = map[string]int{
    "CA": 6234342,
    "TX": 6234343,
    "FL": 8323444,
    "NY": 3324252,
    "PN": 4234235,
    "IL": 2342523,
    "OH": 5234235,
  }
  statePopulations["GE"] = 1222355
  fmt.Println(statePopulations)
  delete(statePopulations, "GE")
  sp := statePopulations // map is reference type
  fmt.Println(sp)
  fmt.Println(sp["GE"]) // 0, pretty confusing
  pop, ok := sp["GE"] // ok should return false
  fmt.Println(pop, ok)
  fmt.Println(sp)
  fmt.Println(statePopulations) // changed the referenced map
}
