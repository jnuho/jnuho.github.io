package main

import (
	// "github.com/jnuho/jnuho.github.io/hellogo/algorithm/structures"
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}

// 2->4->3
// 5->6->4
// 7->0->8

// 9->9->9->9
// 9->9->9->9->9->9->9
// 8->9->9->9->0->0->0->1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t1 := l1
	t2 := l2
	dummyHead := &ListNode{0, nil}
	temp := dummyHead
	n1 := 0
	n2 := 0
	carry := 0
	sum := 0

	for t1 != nil || t2 != nil || carry == 1{

		if t1 != nil {
			n1 = t1.Val
			t1 = t1.Next
		} else {
			n1 = 0
		}
		if t2 != nil {
			n2 = t2.Val
			t2 = t2.Next
		} else {
			n2 = 0
		}

		sum = n1+n2+carry
		carry = sum/10

		temp.Next = &ListNode{sum %10, nil}
		temp = temp.Next
	}

	return dummyHead.Next
}

func iToListNode(i int) *ListNode {
	dummyHead := &ListNode{0, nil}
	temp := dummyHead

	for {
		temp.Next = &ListNode{i%10, nil}
		temp = temp.Next
		i /= 10

		if i == 0 {
			break
		}
	}
	return dummyHead.Next
}

func listNodeToI(l *ListNode) int {
	t := l
	res := 0
	mul := 1

	for t!=nil {
		res += t.Val * mul
		mul *= 10
		t = t.Next
	}
	return res
}

func printNode(l1 *ListNode) {
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next

		if l1 !=nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func main() {
	// 342 + 465 = 807
	l1 := iToListNode(342)
	l2 := iToListNode(465)
	printNode(l1)
	printNode(l2)
	res := addTwoNumbers(l1,l2)
	printNode(res)

	num := listNodeToI(res)
	fmt.Println(num)


	// 9999999 + 9999 = 10009998
	l1 = iToListNode(9999999)
	l2 = iToListNode(9999)
	printNode(addTwoNumbers(l1,l2))
	fmt.Println(listNodeToI(addTwoNumbers(l1,l2)))
}