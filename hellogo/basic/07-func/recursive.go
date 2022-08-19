package main


import (
  "fmt"
)


func factorial(n int) int {
  if n <= 1 {
    return 1
  } else {
    return n * factorial(n-1)
  }
}

func main() {
  n := 5
  fmt.Println(factorial(n))
}
