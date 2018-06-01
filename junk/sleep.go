package main

import (
  "fmt"
  "time"
)

func printAndSleep(x time.Duration) string {
  fmt.Println("INSIDE FUNCTION")
  go func() {
    fmt.Println("Goroutine1")
  }()
  go func() {
    fmt.Println("Goroutine2")
  }()
  go func() {
    fmt.Println("Goroutine3")
  }()
  return msg
}


func main() {
  for {
    fmt.Println("Calling sleep function...")
    result := printAndSleep(time.Second * 5)
    fmt.Println("result: ", result)
  }
}
