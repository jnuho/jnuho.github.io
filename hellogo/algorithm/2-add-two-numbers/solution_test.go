package main


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct{
		name string
		args args
		want *ListNode
	}{
		{
			name: "addTwoNumbers",
			args: args{
				l1: iToListNode(342),
				l2: iToListNode(465),
			},
			want: iToListNode(807),
		},{
			name: "addTwoNumbers",
			args: args{
				l1: iToListNode(0),
				l2: iToListNode(0),
			},
			want: iToListNode(0),
		},{
			name: "addTwoNumbers",
			args: args{
				l1: iToListNode(9999999),
				l2: iToListNode(9999),
			},
			want: iToListNode(10009998),
		},{
			name: "addTwoNumbers",
			args: args{
				l1: iToListNode(10),
				l2: iToListNode(210),
			},
			want: iToListNode(220),
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.want, addTwoNumbers(tt.args.l1, tt.args.l2))
			// expected := tt.want
			// actual := addTwoNumbers(tt.args.l1, tt.args.l2)
			// assert.Equal(t, listNodeToI(expected), listNodeToI(actual))
		})
	}
}