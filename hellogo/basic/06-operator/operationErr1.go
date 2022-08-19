package main

import (
  "fmt"
  "math"
)

const epsilon = 0.000001

func equal(a, b float64) bool {
  diff := math.Abs(a-b)

  if diff <= epsilon {
    return true
  } else {
    return false
  }

}

func main() {
  var a float64 = 0.1
  var b float64 = 0.2
  var c float64 = 0.3

  fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
  fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b,c))

  a= 0.0000000000004
  b= 0.0000000000002
  c= 0.0000000000007

  fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b,c))

}
