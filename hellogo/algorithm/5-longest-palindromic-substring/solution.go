package main

import "fmt"

func Min(i,j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	for i := range s {
		if len(s)-1 -i >= i {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
	}
	return true
}

func longestPalindrome(s string) string {
	max_length := 0
	max_str := ""

	str := ""
	length := 0

	for i := range s {
		length = Min(i, len(s)-1-i)
		for idx:= 0; idx < length; idx++ {
			str = s[i-idx:i+idx+1]
			if isPalindrome(str) {
				if len(str) > max_length {
					max_length = len(str)
					max_str = str
				}
			}
		}
	}
	return max_str
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}