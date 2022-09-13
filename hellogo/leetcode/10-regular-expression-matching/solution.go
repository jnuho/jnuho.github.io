package main

import (
	"fmt"
)

// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
func isMatch(s string, p string) bool {
	for _, r := range p {

		if r == '.' {
			continue
		} else if r == '*' {
			continue
		}
	}
	return false
}

func main() {
	// aa a : does not match the entire string
	// aa a* : * means zero or more of the preceding element
	// ab .* : .* means zero or more (*) of any character (.)
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "b*"))
	fmt.Println(isMatch("ab", ".*"))
}
