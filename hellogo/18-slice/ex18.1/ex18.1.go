package main

import "fmt"

func main() {

	// 슬라이스는 초기화 하지않으면 길이 0 인 슬라이스 만들어짐
	// 길이 초과 해서 슬라이스 멤버에 접근하면  slice[?] 런타임 에러 발생
	// var slice []int
	// if len(slice) == 0 {
	// 	fmt.Println("slice is empty", slice)
	// }
	// ERROR : index out of range
	// slice[1] = 10
	// fmt.Println(slice)


	//WARNING
	// 1 배열선언 - 길이 3 FIXED
	// var array = [...]int{1,2,3}
	// 2 슬라이스 선언
	// var slice0 = []int{1,2,3}


	// 슬라이스 {}로 초기화
	var slice1 = []int{1,2,3}
	var slice2 = []int{1, 5:2, 10:3} // [1 0 0 0 0 2 0 0 0 0 3]

	fmt.Println(slice1)
	fmt.Println(slice2)

	// 슬라이스 make로 초기화
	var slice3 = make([]int, 3)
	slice3[1] = 5
	fmt.Println(slice3, slice3[1])
	
}