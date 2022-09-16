package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	l := 0
	r := len(nums)-1
}

func main() {
	fmt.Println(threeSum([]int{1,2,3}))
}
