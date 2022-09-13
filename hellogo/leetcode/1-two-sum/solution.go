package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	hashmap := map[int]int{}

	// Target = A + B
	for key,A := range nums {
		B := target - A
		_, ok := hashmap[B]

		if ok {
			return []int{hashmap[B], key}
		}
		hashmap[A] = key
	}
	return nil
}


func twoSum(nums []int, target int) []int {
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	nums := []int{1,1}
	target := 2

	res := twoSum(nums, target)
	fmt.Println(res)
}
