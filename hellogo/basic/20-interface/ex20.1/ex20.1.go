package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age int
}

// Student의 String() 메소드
func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {

	p := &Student {"철수", 12}
	s := *p
	fmt.Printf("name=%s, age=%d\n", s.Name, s.Age)

	student := Student{"엘리스", 20}
	// stringer는 Stringer인터페이스이고,
	// Student타입은 String() 메소드를 포함하고 있기 때문에,
	// stringer값으로 student를 대입할 수 있음
	var stringer Stringer = student
	// stringer 인터페이스가 가지고 있는 String() 호출
	// stringer 값으로  Student 타입 student를 가지고 있기 때문에
	// student의 메소드 String() 호출
	fmt.Printf("%s\n",stringer.String())
}
