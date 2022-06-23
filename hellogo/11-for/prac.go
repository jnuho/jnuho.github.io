package main

import (
	"fmt"
)

func main() {
	// 10 9 8 .. 1
	for i := 10; i > 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	// 3-6단 제외 구구단
	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----")
	}

	// 9까지 홀수의 제곱값 출력
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}

}
