package main

import (
  "fmt"
  "strings"
)

func main() {
  x := "../../builds/whatever/tmp"
  y := strings.Split(x, "/")
  r := y[:len(y)-1]
  z := strings.Join(r, "/")
  fmt.Println(z)
}
