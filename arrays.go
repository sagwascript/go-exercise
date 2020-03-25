package main

import "fmt"

func main() {
  // create an array of string with size of 2
  var a [2]string
  a[0] = "Hello"
  a[1] = "world"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  // create an array and give in initial value
  primes := [6]int{2, 3, 5, 7, 11, 13} // items of array can't be more than the length
  fmt.Println(primes)

  evens := [3]int{2, 4} // this ok, the last item will set to 0
  fmt.Println(evens)

  numbers := [...]int{1,2,3,4,5} // let go decide the length based on items we give
  fmt.Println(numbers)

  // array value get copied NOT referenced like in any other languages
  b := a
  b[1] = "guys"
  fmt.Println(b)
  fmt.Println(a)
  fmt.Println(cap(primes))
}
