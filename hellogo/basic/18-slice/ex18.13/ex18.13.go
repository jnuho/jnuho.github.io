package main

import (
	"fmt"
	"sort"
)

// 구조체에 Sort()함수를 이용하기 위해서는 Len(), Less(), Swap() 세가지 메서드 필요!


type Student struct {
	Name string
	Age int
}

// []Student 별칭타입 Students
type Students []Student

func (s Students) Len() int {return len(s)}
func (s Students) Less(i, j int) bool {return s[i].Age < s[j].Age}


// 나이비교
func (s Students) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func main() {
	s := []Student {
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}

	// Students(s): []Student타입을 정렬인터페이스를 포함한 Students로 변환
	// []Student타입은 정렬에 필요한 Len(), Less(), Swap()메서드를 가지고 있지 않으므로
	//	sort.Sort() 함수의 인수로 사용될 수 없음
	sort.Sort(Students(s))
	fmt.Println(s)
}