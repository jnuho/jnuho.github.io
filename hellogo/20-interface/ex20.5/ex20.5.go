package main

import "fmt"

// 2.
// interface{} : 메소드 없는 빈 인터페이스
// 가지고 있어야할 메소드가 하나도 없기때문에 모든 타입이 빈 인터페이스로 쓰일 수 있음
// 어떤값이든 받을 수 있는 함수, 메소드, 변숫값 만들때 사용
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		// 그외 타입인 경우 타입과 값을 출력합니다.
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
