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

	result := make([][]int, 0)

	l, r, mid := 0, len(nums)-1, 0

	for l <= r-2 {
		target := -(arr[l] + arr[r])
		i := binary_search(arr[l+1:r], target)
		if i != -1 {
			result = append(result, []int{arr[l], arr[i], arr[r]})
		} else if target >= 0 {
			l++
		} else if target < 0 {
			r--
		}

	}

	return [][]int{}
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
