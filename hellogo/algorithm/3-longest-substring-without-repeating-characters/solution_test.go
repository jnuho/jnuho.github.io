package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  tests := []struct{
    name string
    args string
    want int
  }{
    {
      name: "lengthOfLongestSubstring",
      args: "aabaab!bb",
      want: 3,
    },
    {
      name: "lengthOfLongestSubstring",
      args: "abba",
      want: 2,
    },
    {
      name: "lengthOfLongestSubstring",
      args: "au",
      want: 2,
    },
    {
      name: "lengthOfLongestSubstring",
      args: "abcabcbb",
      want: 3,
    },
    {
      name: "lengthOfLongestSubstring",
      args: "bbbb",
      want: 1,
    },
    {
      name: "lengthOfLongestSubstring",
      args: "pwwkew",
      want: 3,
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T){
      assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.args))
    })
  }
}