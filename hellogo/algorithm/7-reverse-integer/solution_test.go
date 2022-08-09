package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct{
		name string
		args int
		want int
	}{
		{
			name: "reverse",
			args: 321,
			want: 123,
		},
		{
			name: "reverse",
			args: -123,
			want: -321,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, reverse(tt.args))
		})
	}
}