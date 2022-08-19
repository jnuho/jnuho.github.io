package main

import (
	"fmt"
)

// 3.
// 인터페이스 변환하기
// 구체화된 다른 타입으로 변환하기
type Stringer interface{
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Println(s.String())
	fmt.Printf("Age:%d\n", s.Age)
}


func main() {
	s := &Student{15}
	PrintAge(s)
}