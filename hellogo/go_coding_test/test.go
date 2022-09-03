// sort2. 위에서 아래로
package main

import (
	"fmt"
  "sort"
)

func quick_sort(arr []int, start, end int) {
  if start >= end {
    return
  }
  pivot := start
  left := start+1
  right := end

  for left <= right {
    for left <= end && arr[left] < arr[pivot] {
      left++
    }
    for right > start && arr[right] >= arr[pivot] {
      right--
    }
    if left > right {
      arr[right], arr[pivot] = arr[pivot], arr[right]
    } else {
      arr[left], arr[right] = arr[right], arr[left]
    }
  }
  quick_sort(arr, start, right-1)
  quick_sort(arr, right+1, end)
}

func main() {
  var N,K int
  fmt.Scan(&N, &K)
  arr1 := make([]int, N)
  arr2 := make([]int, N)
  for i:=0; i<N; i++ {
    fmt.Scan(&arr1[i])
  }
  for i:=0; i<N; i++ {
    fmt.Scan(&arr2[i])
  }
  sort.Sort(sort.IntSlice(arr1))
  sort.Sort(sort.IntSlice(arr2))

  for i:=0; i<K; i++ {
    if arr1[i] < arr2[N-i-1] {
      arr1[i], arr2[N-i-1] = arr2[N-i-1], arr1[i]
    }
  }

  sum:=0
  for i:=0; i<N; i++ {
    sum += arr1[i]
  }
  fmt.Println(sum)
}
