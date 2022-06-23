package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(30))

	// int -> 리시버타입에 해당하는 myInt 타입으로 변환 후 메소드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50))
}
 