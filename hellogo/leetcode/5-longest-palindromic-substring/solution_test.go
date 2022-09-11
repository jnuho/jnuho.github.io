package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "longestPalindrome",
			args: "babad",
			want: "bab",
		},
		{
			name: "longestPalindrome",
			args: "cbbd",
			want: "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPalindrome(tt.args))
		})
	}
}