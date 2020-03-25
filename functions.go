package main

import (
  "fmt"
)

func main() {
  sayMessage("Hello", "Stacey", func (v string) { // pass function as argument
    fmt.Println(v)
  })
  res := sum(1,2,3,4,5)
  fmt.Println(*res)
  d, err := divide(5.0, 0.0) // don't forget to retrieve the multiple return
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(d)
}

func sayMessage(greeting, name string, f func(v string)) { // accept function as argument
  // fmt.Println(greeting, name)
  f(greeting+name)
}

// variadic function
func sum(values ...int) *int { // take all the arguments, looks like spread in js, it has to be the last
  fmt.Println(values)
  result := 0
  for _, v := range values {
    result += v
  }
  return &result
}

func divide(a, b float64) (float64, error) { // return multiple value
  if b == 0.0 {
    return 0.0, fmt.Errorf("Cannot divide by zero")
  }
  return a / b, nil
}
