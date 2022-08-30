package main

import (
	"fmt"
)

// O(N^2)
func selection_sort(arr []int) {
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
}

// O(N^2)
func insert_sort(arr []int) {
	for i := 0; i<len(arr); i++ {
		for j := i; j>=1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// O(NlogN)
func quick_sort(arr []int, ) {
	pivot := arr[0]
}


func main() {
	arr := []int {7,5,9,0,3,1,6,2,4,8}
	selection_sort(arr)
	fmt.Println(arr)

	arr = []int {7,5,9,0,3,1,6,2,4,8}
	insert_sort(arr)
	fmt.Println(arr)
}