package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1==p2)
	fmt.Printf("p2 == p3 : %v\n", p2==p3)


	// 포인터의 기본값 nil
	var p *int
	if p != nil {
		fmt.Println("포인터 p는 유효한 메모리 공간을 가리킴")
	} else {
		fmt.Println("포인터 p == nil 어떤 메모리 공간도 가리키고 있지 않음")
	}

	// 포인터는 변수대입 또는 함수 인수전달 시 메모리 공간을 적게 사용하면서 사용 가능
}