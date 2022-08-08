package main

import (
	"fmt"
	"strconv"
	"math"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func reverse(x int) int {
	ux := 0
	sign := 1
	if x < 0 {
		ux = -x
		sign = -1
	} else {
		ux = x
	}

	s := []rune(strconv.Itoa(ux))

	for i,j := 0, len(s)-1; i<j; i,j = i+1,j-1 {
		s[i], s[j] = s[j], s[i]
	}
	res, err := strconv.Atoi(string(s))
	if err !=nil {
		fmt.Println(err)
	}

	if sign*res < - powInt(2,31) || sign*res > powInt(2,31)-1 {
		return 0
	}
	return sign *res
}

func main() {
	fmt.Println(reverse(321))
	fmt.Println(reverse(-123))
}