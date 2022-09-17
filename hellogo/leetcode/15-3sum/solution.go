package main

import (
	"fmt"
	//"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)

	l := 0
	r := len(nums)-1
	for l < r-1 {

	}
}

// returns array index of a target element
// returns -1 if not found
func binary_search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	l, r, mid := 0, len(arr)-1, 0

	for l <= r {
		mid = (l+r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return -1
}

func main() {
	arr := []int {1,2,3}
	fmt.Println(threeSum([]int{1,2,3}))
}
