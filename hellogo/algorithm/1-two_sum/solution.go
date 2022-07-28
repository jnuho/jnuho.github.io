package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) <2 {
		return []int{}
	}

	numToIndex := make(map[int]int)
	for idx, num := range nums {
		need := target - num

		mIdx, ok := numToIndex[need]
		if ok {
			return []int{mIdx, idx}
		}

		numToIndex[num] = idx
	}

	return nil
}

/**
Assume 'nums' is a sorted array in ascending order
*/
func twoSumSorted(nums []int, target int) []int {
	front := 0
	rear := len(nums) -1

	for front < rear {
		need := nums[front] + nums[rear]
		if need == target {
			return []int{front, rear}
		} else if need < target {
			front = front +1
		} else  {
			rear = rear-1
		}
	}
	return nil
}

func main() {
	nums := []int{5,3,1,2,4}
	target := 5
}