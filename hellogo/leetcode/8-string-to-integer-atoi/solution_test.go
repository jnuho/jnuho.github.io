package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// func myAtoi(s string) int {

func Test(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "myAtoi",
			args: "42",
			want: 42,
		},{
			name: "myAtoi",
			args: "   -42",
			want: -42,
		},{
			name: "myAtoi",
			args: "4193 with words",
			want: 4193,
		},{
			name: "myAtoi",
			args: "starts with words 123123",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, myAtoi(tt.args))
		})
	}
}