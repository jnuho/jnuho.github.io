package main

import _ "fmt"

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	if r,ok := stringer.(Reader); ok {
		r.Read()
	}
}

func main() {
}