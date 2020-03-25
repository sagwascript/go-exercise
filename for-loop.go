package main

import "fmt"

func main() {
  sum := 0
  for i := 1; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
    fmt.Println(i, j)
  }

  // omit init and post statements
  add := 1
  for ; add < 16; {
    fmt.Println(add)
    add += add
  }
  fmt.Println(add)

  // while-like for loop
  iter := 1
  for iter < 3 {
    fmt.Println(iter)
    iter += 1
  }
  
  // DON'T, this will create an infinite loop
  // for {

  // }

  Loop: // label
  for i := 1; i <= 3; i++ {
    for j := 1; j <=3; j++ {
      fmt.Println(i * j)
      if i * j >= 3 {
        // break // this kind of break will only break the inner loop
        break Loop // inner loop can break the outer loop
      }
    }
  }
}
