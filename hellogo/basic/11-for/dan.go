package main

import (
	"fmt"
)

func main() {
	dan := 2
	b := 1

	for {
		for {
			fmt.Printf("%d x %d = %d \n", dan, b, dan*b)

			b++

			if b == 10 {
				break
			}
		}
		b = 1
		dan++
		fmt.Println("=====")
		if dan == 10 {
			break
		}
	}
}
