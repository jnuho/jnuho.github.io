package main

import "fmt"

// func add(a, b int) int {
// 	return a+b
// }

// func mul(a, b int) int {
// 	return a*b
// }

// 함수리터럴
type opFunc func (int, int) int

// 반환 타입 : func (int, int) int
func getOperator(op string) opFunc {
	if op == "+" {
		// return add
		return func(a, b int) int {
			return a+b
		}
	} else if op == "*" {
		// return mul
		return func(a, b int) int {
			return a*b
		}
	} else { // +나 *가 아니면 nil 반환
		return nil
	}

}

func main() {
	// var operator func (int, int) int
	// operator = getOperator("*")

	// var result = operator(3, 4)
	// fmt.Println(result)

	operator := getOperator("*")
	result := operator(3, 4)
	fmt.Println(result)


	// 리터럴 함수 호출
	// 1
	fn := func(a,b int) int {
		return a +b
	}
	res := fn(3,4)
	fmt.Println(res)

	// 2
	res = func(a,b int) int {
		return a +b
	}(3,4)
	fmt.Println(res)
}
