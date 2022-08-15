package main

import (
	"fmt"
	"math"
)
func powInt(x,y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// 12345
func reverse(x int) int {

	var result int

	for x != 0 {
		result = result*10 + x%10
		if result < - powInt(2,31) || result > powInt(2,31)-1 {
			return 0
		}
		x /= 10
	}

	return result
}

func main() {
	fmt.Println(reverse(321))
	fmt.Println(reverse(-123))
}