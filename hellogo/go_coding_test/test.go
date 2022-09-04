package main

import (
  "fmt"
)

var arr []int

func init() {
  arr = make([]int, 100)
}

func fibo(n int) int {
  if n==1 || n == 2 {
    return 1
  }
  if arr[n] != 0 {
    return arr[n]
  }
  arr[n] = fibo(n-1) + fibo(n-2)
  return arr[n]
}

func main() {
  fmt.Println(fibo(90))
}
