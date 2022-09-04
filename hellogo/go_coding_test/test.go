// 2.떡볶이 떡 만들기 - 이진탐색
// 4 6
// 19 15 10 17
// 15
package main

import (
  "fmt"
)

func getMax(arr []int) int {
  max :=arr[0]
  for i:=0; i<len(arr); i++ {
    if arr[i] > max {
      max = arr[i]
    }
  }
  return max
}
func getDduck(arr []int, cutter int) int {
  result := 0
  for i:=0; i<len(arr); i++ {
    if arr[i] > cutter {
      result += arr[i]-cutter
    }
  }
  return result
}

func main() {
  var N,M int
  fmt.Scan(&N, &M)

  arr := make([]int, N)
  for i:=0; i<N; i++ {
    fmt.Scanf("%d", &arr[i])
  }
  max := getMax(arr)
  left, right := 0, max
  var mid, sum int

  for left <=right {
    mid = (left+right) / 2
    sum = getDduck(arr,mid)
    if sum == M {
      fmt.Println(mid)
      break
    } else if sum > M {
      left = mid+1
    } else {
      right = mid-1
    }
  }
}
