package main

import (
	"math"
	"strconv"
	"fmt"
)

func powInt(x, n int) int {
	return int(math.Pow(float64(x), float64(n)))
}

// "42" -> 42
// "   -42" -> -42
// "4193 with words" -> 4193
// "0032" -> 32
func myAtoi(s string) int {
	// leading whitespace
	// - +
	// num
	sign := 1

	for i,r := range s {
		if string(r) == " " {
			continue
		} else if string(r) == "-" {
			sign = -1
			s = string(s[i+1:])
			break
		} else if string(r) == "+" {
			s = string(s[i+1:])
			break
		} else if r >= '0' && r <= '9' {
			s = string(s[i:])
			break
		} else {
			return 0
		}
	}

	for i,r := range s {
		if r < '0' || r > '9' {
			s = string(s[:i])
			break
		}
	}
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int!!!")
	}
	result = sign * result
	
	if result < -powInt(2, 31) {
		return -powInt(2,31)
	} else if result > powInt(2,31) -1 {
		return powInt(2,31)-1
	}
	return result
}
