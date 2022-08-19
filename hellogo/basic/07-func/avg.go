package main

import (
  "fmt"
)

func Avg(a int, b int)  int {
  sum := a+b
  return sum / 2
}

func main() {
  c := Avg(3,6)

  fmt.Println(c)
}
