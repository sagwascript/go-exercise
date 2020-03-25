package main

import "fmt"

func main() {
  num := 24
  pointerNum := &num // reference to pointer of variable num
  num = 41 // will also change value accessed by pointer

  fmt.Println(*pointerNum) // read from pointer's value, see the asterisk
  fmt.Println(pointerNum) // return the memory address

  *pointerNum = 7 // change from pointer
  fmt.Println(*pointerNum) 
  fmt.Println(num)

  var a int = 42
  var b *int = &a // declare a pointer, & is called addressof operator
  a = 23
  fmt.Println(a, *b) // asterisk is for dereferencing pointer
  *b = 44
  fmt.Println(a, *b)

  var ms *myStruct
  fmt.Println(ms) // zero value of pointer is <nil>
  ms = new(myStruct)
  // (*ms).foo = 42 // you can do this
  ms.foo = 42
  // fmt.Println((*ms).foo) // you can do this too
  fmt.Println(ms.foo) // but this just works
}

type myStruct struct {
  foo int
}
