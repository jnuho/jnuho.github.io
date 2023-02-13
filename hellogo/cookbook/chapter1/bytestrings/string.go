package bytestrings

import (
	"fmt"
	"strings"
)

func SearchString() {
	s := "this is a test"

	// contains "this"
	fmt.Println(strings.Contains(s, "this"))

	// contains any of a,b,c
	fmt.Println(strings.ContainsAny(s, "abc"))
}
