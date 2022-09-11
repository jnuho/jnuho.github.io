package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)

	sort.Slice(nums, func(i,j int) bool {
		return nums[i] < nums[j]
	})
	// 1 2 3 4 5
	// 1 2 3 4
	length := len(nums)
	if length %2 == 1 {
		return float64 (nums[length/2])
	} else {
		return float64 (nums[length/2 -1] + nums[length/2])/2.0
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}