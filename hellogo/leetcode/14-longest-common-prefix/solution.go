package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	str := strs[0]
	pos :=0
	for i:=0; i< len(str)+1; i++ {
		for j:=1; j <len(strs); j++ {
			if len(strs[j]) < i || !strings.Contains(strs[j][:i], str[:i]) {
				return str[:i-1]
			}
		}
		pos = i
	}

	return str[:pos]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}
