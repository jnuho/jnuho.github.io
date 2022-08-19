package main

import "fmt"

type account struct {
	balance int
	firstName string
	lastName string
}


// 포인터 메소드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메소드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값 반환하는 값 타입 메소드

func (a3 account) withdrawReturnValue (amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park" }
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	// 값이 안변함!
	// mainA와 a2변수는 서로 다른 메모리 주소를 가지게 됨 (서로 다른 인스턴스)
	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20) // mainA,mainB,a3 모두 다른 메모리
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}