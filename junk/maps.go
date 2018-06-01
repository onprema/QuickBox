package main

import (
  "fmt"
)

func main() {
  m := make(map[string]int)
  m["lee"] = 5
  fmt.Printf("%v\n", m)
  fmt.Println(m)
}
