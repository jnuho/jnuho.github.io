package main

import "fmt"

func main() {
	// len=3, cap=3 슬라이스
	slice1 := []int{1,2,3}

	// cap=3이므로 용량 초과! : 기존배열이 아닌 신규 배열 생성하여 따로 바라보도록 됨
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("AFTER change second element")
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("AFTER append 500")
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))
}