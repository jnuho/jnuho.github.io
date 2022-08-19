package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}

	for i:=0; i< 10; i++ {
		fmt.Println("라라라")
	}

	// 초기문 생략
	// for ; 조건문  ; 후처리 {}

	// 후처리 생략
	// for 초기문; 조건문; {}

	// 조건문만 있는 경우
	// for ;조건문; {}
	// for 조건문 {}

	// 무한루프
	// for true {}
	// for {}
	count := 0
	for count < 10 {
		fmt.Printf("%d ", count+1)
		count++
	}
	fmt.Println()

	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

}
