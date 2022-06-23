package main

import (
  "fmt"
)

func main(){
  var x int32 = 7
  //var y float32 = 3.14

  fmt.Println(x + 1)
  fmt.Println(x - 1)
  fmt.Println(x * 1)
  fmt.Println(x / 1)
  fmt.Println(x % 1)

  fmt.Println(x & 1)
  fmt.Println(x | 1)
  fmt.Println(x ^ 1)
  fmt.Println(x &^ 1)

  fmt.Println(x << 1)
  fmt.Println(x >> 1)
}
