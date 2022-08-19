package main

type Stringer interface {
	String() string
}

type Student struct {

}

func main() {
	var stringer Stringer
	stringer.(*Student) // ERROR : Student는 String()을 구현하고 있지 않음
}