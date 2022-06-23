package main

import "fmt"

/**
특정인덱스 위치에 요소 추가하기
*/
func main() {
	slice := []int{1,2,3,4,5,6}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 0) // len:7, cap:12 (=6*2)
	fmt.Println(slice, len(slice), cap(slice))

	idx := 2
	for i:=len(slice)-1;  i>idx; i-- {
		slice[i] = slice[i-1]
	}
	slice[idx] = 100
	fmt.Println(slice, len(slice), cap(slice))


	// apend()로 코드 개선하기
	// [100,3,4,5,6]은 임시 슬라이스이며 불필요한 메모리를 사용
	slice = []int{1,2,3,4,5,6}
	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	fmt.Println(slice, len(slice), cap(slice))


	// 불필요한 메모리사용 없도록 개선
	slice = []int{1,2,3,4,5,6}
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100
	fmt.Println(slice, len(slice), cap(slice))
}