package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



// func isMatch(s string, p string) bool {
func Test(t *testing.T) {
	type args struct  {
		s string // input string
		pattern string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isMatch",
			args: args{
				s: "aa",
				pattern: "a",
			},
			want: false,
		},{
			name: "isMatch",
			args: args{
				s: "aa",
				pattern: "a*",
			},
			want: true,
		},{
			name: "isMatch",
			args: args{
				s: "aa",
				pattern: ".*",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, isMatch(tt.args.s, tt.args.pattern))
		})
	}
}