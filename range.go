package main

import "fmt"

func main() {
  pow := [5]int{ 5, 1, 3, 2, 4 }
  // loop through key value pair
  for index, value := range pow {
    fmt.Println(index, ": ", value)
  }

  // get index only without value
  for index := range pow {
    fmt.Println("index: ", index)
  }

  // get value only without index
  for _, value := range pow {
    fmt.Println("item: ", value)
  }

  // use range in a string
  s := "hello Go!"
  for k, v := range s {
    fmt.Println(k, string(v)) // cast it back to string since it return a byte representing the ascii
  }
}
