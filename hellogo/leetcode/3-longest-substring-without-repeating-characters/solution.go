package main

import "fmt"

func Max(x,y int) int {
  if x > y {
    return x
  } else {
    return y
  }
}

func lengthOfLongestSubstring(s string) int {
  if len(s) < 2 {
    return len(s)
  }
  // {문자: 인덱스} 저장
  mymap := map[rune]int{}

  length := 0
  max_length := 0
  // sliding window
  // start_i := 0

  // 문자 1개씩 검색
  for i, c := range s {
    // 맵에 포함여부에 따라 최대길이 업데이트
    if _, ok := mymap[c]; !ok {
      length++
    } else {
      mymap_c := mymap[c]
      length = i - mymap_c
      for k, v := range mymap {
        if v <= mymap_c {
          delete(mymap, k)
        }
      }
    }
    max_length = Max(length, max_length)
    mymap[c] = i
  }

  return max_length
}

func main() {
  fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
  fmt.Println(lengthOfLongestSubstring("abba"))
  fmt.Println(lengthOfLongestSubstring("au"))
  fmt.Println(lengthOfLongestSubstring("abcabcbb"))
  fmt.Println(lengthOfLongestSubstring("bbbbb"))
  fmt.Println(lengthOfLongestSubstring("pwwkew"))
}