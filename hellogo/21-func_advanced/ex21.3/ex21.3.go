package main

import "fmt"

func add(a, b int) int {
	return a+b
}

func mul(a, b int) int {
	return a*b
}

// ALIAS 지정 가능
type opFunc func (int, int) int
// func getOperator(op string) opFunc {...}

// 반환 타입 : func (int, int) int
// func getOperator(op string) func (int, int) int {
func getOperator(op string) opFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
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
}