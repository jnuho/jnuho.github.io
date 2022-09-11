package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := [] struct {
		name string
		args int
		want bool
	}{
		{
			name: "isPalindrome",
			args: 121,
			want: true,
		},{
			name: "isPalindrome",
			args: -121,
			want: false,
		},{
			name: "isPalindrome",
			args: 10,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, isPalindrome(tt.args))
		})
	}
}