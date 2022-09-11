package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct{
		s string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "convert",
			args: args{
				"PAYPALISHIRING",
				3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "convert",
			args: args{
				"PAYPALISHIRING",
				4,
			},
			want: "PINALSIGYAHRPI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, convert(tt.args.s, tt.args.numRows))
		})
	}

}