package main

import "fmt"


func main() {
	var slice = []int{10, 20, 30}

	for i :=0; i <len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i,v := range slice {
		fmt.Println(i, v)
	}
}