package main

import (
  "fmt"
  "math/big"
)


func main() {
  a, _ :=  new(big.Float).SetString("0.1") // a
  b, _ :=  new(big.Float).SetString("0.2") // b
  c, _ :=  new(big.Float).SetString("0.3") // c

  d := new(big.Float).Add(a,b) // a+b
  fmt.Println(a, b, c, d)
  fmt.Println(c.Cmp(d)) // 0: equals (정밀도 높음)
}
