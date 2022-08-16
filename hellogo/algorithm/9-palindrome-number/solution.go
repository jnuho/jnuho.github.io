package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	rx := 0
	for x != 0 {
		rx = rx * 10 + x % 10
		x /= 10
	}

	return tmp == rx
}

func main() {
	fmt.Println()
}