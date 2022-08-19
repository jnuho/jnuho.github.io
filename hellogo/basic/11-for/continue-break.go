package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 || i == 6 {
			continue
		}

		fmt.Println("7 * ", i, "=", 7*i)
	}
}
