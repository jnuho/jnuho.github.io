// 10 7
// 1 3 5 7 9 11 13 15 17 19
// 4
package main

import (
	"fmt"
)

func binary_search(arr []int, target, start, end int) int {
  if start > end {
    return -1
  }

  if len(arr) < 1 {
    return -1
  }

  mid := (start+end) / 2
  if arr[mid] == target {
    return mid
  } else if arr[mid] > target {
    return binary_search(arr, target, start, mid-1)
  } else {
    return binary_search(arr, target, mid+1, end)
  }
}

func main() {
  var N int // 리스트 개수
  var target int // 찾는 숫자
  fmt.Scan(&N, &target)

  arr := make([]int, N)
  for i:=0; i<N; i++ {
    fmt.Scan(&arr[i])
  }

  result := binary_search(arr, target, 0, len(arr)-1)
  if result == -1 {
    fmt.Println("NOT FOUND!")
  } else {
    fmt.Println(result + 1)
  }
}

