package main

import (
	"fmt"
)

// O(N^2)
func selection_sort(arr []int) []int {
	var min_idx int
	for i:=0; i<len(arr); i++ {
		min_idx = i
		for j:=i+1; j<len(arr); j++ {
			if arr[min_idx] > arr[j] {
				min_idx = j
			}
		}
		arr[min_idx], arr[i] = arr[i], arr[min_idx]
	}
	return arr
}

// O(N^2)
func insert_sort(arr []int) []int {
	for i := 0; i<len(arr); i++ {
		for j := i; j>=1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// O(NlogN)
func quick_sort(arr []int, start, end int) {
	if start >=end {
		return
	}
	pivot := start

	left := start+1
	right := end

	for left <= right {
		for left <= end && arr[left] <= arr[pivot] {
			left++
		}
		for right > start && arr[right] > arr[pivot] {
			right--
		}

		if right > left {
			arr[left], arr[right] = arr[right], arr[left]
		} else {
			// 엇갈림 -> for-loop exit
			arr[right], arr[pivot] = arr[pivot], arr[right]
		}
	}
	quick_sort(arr, start, right-1)
	quick_sort(arr, right+1, end)
}


func counting_sort(arr []int) []int{

}

func main() {
	arr := []int {7,5,9,0,3,1,6,2,4,8}
	fmt.Println(selection_sort(arr))

	arr = []int {7,5,9,0,3,1,6,2,4,8}
	fmt.Println(insert_sort(arr))

	arr = []int {7,5,9,0,3,1,6,2,4,8}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int {7,5,9,0,3,1,6,2,4,8}
	fmt.Println(counting_sort(arr))
}
