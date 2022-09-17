package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct{
		name string
		args []int
		want [][]int
	}{
		{
			name: "threeSum",
			args: []int{-1,0,1,2,-1,-4},
			want: [][]int{[]int{-1,0,1}, []int{-1,-1,2}},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, threeSum(tt.args))
		})
	}

}


