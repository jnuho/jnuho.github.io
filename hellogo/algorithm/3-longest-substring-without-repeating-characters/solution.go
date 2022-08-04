package main

import (
  "fmt"
  "math"
)

func lengthOfLongestSubstring(s string) int {
  if len(s) < 2 {
    return len(s)
  }

  mymap := map[rune]int{}
  length := 0
  max_length := 0
  for i, c := range s {
    if _, ok := mymap[c]; !ok {
      length++
    } else {
      max_length = int(math.Max(length, max_length))
      length -= (mymap[c]+1)
    }
    mymap[c] = i
  }

  return max_length
}

func main() {
  fmt.Println(lengthOfLongestSubstring("abcabcbb"))
  fmt.Println(lengthOfLongestSubstring("bbbbb"))
  fmt.Println(lengthOfLongestSubstring("pwwkew"))
}