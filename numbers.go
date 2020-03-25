package main

import (
  "fmt"
)

// beware const can shadow
const c int16 = 23

func main() {
  // this shadow the package level const
  const c int8 = 127
  var name string = "john"
  var byteName []byte = []byte(name)
  // complex64 used float32 real and float32 imaginary
  var res complex64 = complex(5, 12)
  // complex64 used float64 real and float64 imaginary
  var other complex128 = complex(2, 3)
  a := real(res) // get real part of complex number
  b := imag(res) // get imaginary part of complex number
  fmt.Printf("%v\n", res)
  fmt.Printf("%v %T\n", a, a)
  fmt.Printf("%v %T\n", b, b)
  fmt.Printf("%v %T\n", other, other)
  fmt.Printf("%v %T\n", real(other), real(other))
  fmt.Printf("%v %T\n", imag(other), imag(other))
  fmt.Println(byteName)
  fmt.Println(c)
}
