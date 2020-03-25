package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

type Doctor struct { // exported because the first capital letter
  number int // you can change this first letter to uppercase to make it accessible through other package
  actorName string
  companions []string
}

func main() {
  coord := Vertex{3, 4}
  // access struct field
  x := coord.X
  y := coord.Y
  fmt.Println(x)
  fmt.Println(y)

  coord2 := Vertex{7, 2}
  p := &coord2
  p.X = 9 // reassign struct field using pointer without dereferencing
  fmt.Println(coord2)

  coord3 := Vertex{X: 8} // set X only, Y is set to 0, default value of int
  fmt.Println(coord3)

  // coord4 := Vertex{12} // won't work, expect two values
  // fmt.Println(coord4)

  aDoctor := Doctor{
    number: 3,
    actorName: "Jon Pertwee",
    companions: []string {
      "Liz Shaw",
      "Jo Grant",
      "Sarah Jane Smith",
    },
  }
  doctor2 := Doctor{ // not recommended, cause it has to be in correct order and complete
    3, "John", []string {"Mark", "Matthew"},
  }
  fmt.Println(aDoctor.companions[1])
  fmt.Println(doctor2)

  // anonymous struct
  doctor3 := struct{name string}{name: "Luke"}
  fmt.Println(doctor3)

  // struct is value type, so it copied
  // use pointer to reference the same struct variable
  doctor4 := aDoctor
  doctor4.actorName = "Alejandro"
  fmt.Println(aDoctor)
  fmt.Println(doctor4)
}
