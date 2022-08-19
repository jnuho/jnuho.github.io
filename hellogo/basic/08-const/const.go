package main

import (
	"fmt"
)

const Pig int = 0
const Cow int = 1
const Chicken int = 2

const PI = 3.14
const PI2 float64 = 3.14

// iota로 간편하게 열거값 사용 하기
// 1씩증가. 소괄호 벗어나면 다시 초기화 됨
const (
	Red   int = iota // 0
	Blue  int = iota // 1
	Green int = iota // 2
)

// const를 소괄호()로 묶고 iota 사용하면
// 0부터 1씩 차례로 증가하며 값이 초기화 됨
const (
	C1 uint = iota + 1 // 1 = 0 + 1
	C2                 // 2 = 1 + 1
	C3                 // 3 = 2 + 1
)

const (
	BitFlag1 uint = iota + 1 // 1 = 1 << 0
	BitFlag2                 // 2 = 1 << 1
	BitFlag3                 // 4 = 1 << 2
	BitFlag4                 // 8 = 1 << 3
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	}

}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	var a int = PI * 100
	// var b int = PI2 * 100 // ERROR!
	fmt.Println(a)
	// fmt.Println(b)

	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)

	fmt.Println(BitFlag1)
	fmt.Println(BitFlag2)
	fmt.Println(BitFlag3)
}


