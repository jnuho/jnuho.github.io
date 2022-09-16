package main

import (
	"fmt"
)

func min(i,j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i,j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	area := 0
	l := 0
	r := len(height) -1
	for l < r {
		newarea := (r-l) * min(height[r],height[l])
		area = max(area, newarea)
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return area
}

func main() {
	height := []int {1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

