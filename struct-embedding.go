package main

import (
  "fmt"
)

type Animal struct {
  Name string
  Origin string
}

type Bird struct {
  Animal // embed other struct, it COMPOSE!!
  SpeedKPH float32
  CanFly bool
}

func main() {
  b := Bird{}
  b.Name = "emu" // get it from embedded struct
  b.Origin = "Australia"
  b.SpeedKPH = 48
  b.CanFly = false
  fmt.Println(b)
  // using literal syntax
  c := Bird{
    Animal: Animal{Name: "Slo", Origin: "Australia"}, // should type it explicitly
    SpeedKPH: 48,
    CanFly: false,
  }
  fmt.Println(c)
}
