package main

import (
	"fmt"
)

var d []int

func initD(N int){
  d = make([]int, N)
}

func max(i,j int ) int {
  if i>j {
    return i
  } else {
    return j
  }
}

func main() {
  var N int
  fmt.Scanln(&N)
  initD(N)

  arr := make([]int, N)
  for i:=0; i<N; i++ {
    fmt.Scanf("%d", &arr[i])
  }

  d[0] = arr[0]
  d[1] = max(arr[0], arr[1])

  for i:=2; i< len(d); i++ {
    d[i] = max(d[i-1], d[i-2] + arr[i])
  }
  fmt.Println(d[N-1])
}
