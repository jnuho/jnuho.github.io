package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4,5}
	slice2 := make([]int, 3, 10) // len:3, cap:10
	slice3 := make([]int, 10) // len:10, cap:10

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println("slice2: ", cnt1, slice2)
	fmt.Println("slice3: ", cnt2, slice3)

}