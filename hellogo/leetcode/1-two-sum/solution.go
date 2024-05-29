package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if find, ok := hashmap[target-nums[i]]; ok {
			return []int{find, i}
		}
		hashmap[nums[i]] = i
	}
	return []int{}
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	nums := []int{1, 1}
	target := 2

	res := twoSum(nums, target)
	fmt.Println(res)
}
