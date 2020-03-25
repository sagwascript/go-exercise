package main

import (
  "net/http"
)

func main() {
  // this will serve response in localhost
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Go"))
  })
  http.HandleFunc("/guys", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Guys"))
  })
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    panic(err.Error()) // should be panic when address is already used for example
  }
}
