









package main

import (
	"fmt"
)

type Comparable interface {
	LessThan(other interface{}) bool
}

func Max[T Comparable](nums []T) T {
	max := nums[0]
	for _, n := range nums {
		if n.LessThan(max) {
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	nums := []int{1,2,3,4,5}
	max := Max(nums)
	fmt.Println(max)
}




