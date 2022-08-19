package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}



func main() {
	userPointer := NewUser("AAA", 23)
	// NewUser 함수 내에서 로컬 선언된 변수 u는
	// 함수 종료후에도 남아있음
	// : Go언어 탈출검사를 통해
	// u변수의 인스턴스가 함수외부로 공개된것을 분석해내서
	// u변수를 스택메모리가 아닌 힙메모리에 할당함
	// : no dangling error
	스택메모리가 아닌 힙메모리에
	fmt.Println(userPointer)
	fmt.Printf("userPointer=%p\n", userPointer)
}