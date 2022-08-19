package main

import "fmt"

func main() {
	
	var str string = "Hello World"
	str = "How are you"
	// ERROR: string is immutable
	// str[2] = 'a'

	fmt.Println(str)

	var slice []byte = []byte(str)
	slice[2] = 'a'
	fmt.Printf("%s\n", slice)

	// str = string(slice)
	// fmt.Printf("%s\n", str)
}