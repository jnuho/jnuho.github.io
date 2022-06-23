package main

import (
	"fmt"
)

// label사용은 지양 - 혼동 올 수 있음
func find45(a int) (int, bool) {
	for b := 1; b < 10; b++ {
		if b*a == 45 {
			return b, true
		}
	}
	return 0, false
}



func main() {
	a := 1
	b := 0
	found := false

	for ; a <= 9; a++ {
		if b, found = find45(a); found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
