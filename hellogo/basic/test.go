package main

import (
	"fmt"
)

func main() {
	a := [5]int{1,2,3,4,5}
	b := [5]int{500,400,300,200,100}

	// a
	for i,v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	// b
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	// a가 복사된 b
	fmt.Println()
	for i,v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

