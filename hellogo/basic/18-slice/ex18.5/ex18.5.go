package main

import "fmt"

func PrintSlice(name string, slice []int) {
	fmt.Println(name, slice, "len: ", len(slice), "cap: ", cap(slice))
}

func main() {
	// len = 3, cap = 5
	// cap-len = 2
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	PrintSlice("slice1", slice1)
	PrintSlice("slice2", slice2)

	// slice1, slice2 같은 배열을 가리키므로 '둘다' 변경됨
	// 다만 cap, len 값은 달라서 조회 되는 slice는 차이 있음
	slice1[1] = 100

	PrintSlice("slice1", slice1)
	PrintSlice("slice2", slice2)

	slice1 = append(slice1, 500)

	PrintSlice("slice1", slice1)
	PrintSlice("slice2", slice2)
}