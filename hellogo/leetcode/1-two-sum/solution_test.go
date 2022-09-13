package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	type args struct {
		nums []int
		target int
	}

	tests := []struct{
		name string
		args args
		want []int
	}{
		{
			name: "twoSum",
			args: args{
				nums: []int{2,7,11,15},
				target: 9,
			},
			want: []int{0,1},
		},
		{
			name: "twoSum",
			args: args{
				nums: []int{3,2,4},
				target: 6,
			},
			want: []int{1,2},
		},
		{
			name: "twoSum",
			args: args{
				nums: []int{3,3},
				target: 6,
			},
			want: []int{0,1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}

}
