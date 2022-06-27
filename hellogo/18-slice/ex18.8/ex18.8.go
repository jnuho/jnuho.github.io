package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4,5}
	// slice2 := make([]int, len(slice1))
	slice2 := append([]int{}, slice1...)
	// slice2 := append([]int{}, slice1[0], slice1[1], slice1[2], slice1[3], slice1[4])

	for i,v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

}