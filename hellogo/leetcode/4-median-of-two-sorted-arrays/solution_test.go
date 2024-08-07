package main

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "findMedianSortedArrays",
			args: args {
				nums1: []int{1,3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "findMedianSortedArrays",
			args: args {
				nums1: []int{1,2},
				nums2: []int{3,4},
			},
			want: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			assert.Equal(t, tt.want, findMedianSortedArrays(tt.args.nums1, tt.args.nums2))
		})
	}
}