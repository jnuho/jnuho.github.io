package main

import (
	"fmt"
	"strings"
)

/**
string내, 알파벳 반복횟수 계산하기
결과가 -1 이면 prime number
*/
func findDivNum(N int) int {
	res := N
	for i := 2; i <= N; i++ {
		if N%i == 0 {
			if i <= res {
				res = i
			}
		}
	}
	if res == N {
		return -1
	} else {
		return res
	}
}

func main() {
	fmt.Println(findDivNum(967))
	a := strings.Repeat("A", 5)
	fmt.Println(a)
}
