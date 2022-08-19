package main

import (
  "fmt"
)

func Add(a int, b int)  int {
  sum := a+b
  return sum
}

func main() {
  c := Add(3,6)
  fmt.Println(c)
}
