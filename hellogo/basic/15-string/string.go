package main

import "fmt"

func main() {
	s1 := "Hello string variable!"
	s2 := `Special \t characters \t\t \n inside backticks are ignored\n
	backticks also ignores new lines
	line3
	line4
	line5`

	fmt.Println(s1)
	fmt.Println(s2)

}