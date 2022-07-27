package main

import (
	"fmt"
)

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// my_map[2]=0
// my_map[7]=1
// my_map[7]=1
// my_map[11]=2
// my_map[15]=3

// target = 6
// my_map[3]=0
// my_map[2]=1
// my_map[4]=2

func TwoSums(nums []int, target int) []int {
	my_map := make(map[int]int)

	for i, val := range nums {
		fmt.Println(i, val)
		my_map[val] = i
		if j, ok  := my_map[target - val]; ok && i != j{
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	// target := 9
	// nums := []int{2, 7, 11, 15}
	target := 6
	nums := []int{3, 2, 4}
	res := TwoSums(nums, target)

	fmt.Printf("len(res)= %d\n", len(res))
	fmt.Println(res)
}