package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) printListNode() {
	if n != nil {
		fmt.Printf("%d", n.Val)
	}
	if n.Next != nil {
		fmt.Printf("->")
		n.Next.printListNode()
	} else {
		fmt.Println()
	}
}

//321 1->2->3
func makeListNode(num int) *ListNode {
	l := &ListNode{}
	dummy := l
	for num > 0 {
		// Val
		l.Val = num % 10

		// Next
		num = num / 10
		if num > 0 {
			l.Next = &ListNode{}
			l = l.Next
		}

	}
	return dummy
}
func (n *ListNode) listNodeToNum() int {
	t := n
	dec := 1
	num := 0

	for t != nil {
		num += dec * t.Val

		t = t.Next
		dec *= 10
	}
	return num
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//t1 = l1
	//t2 = l2
	//dummyHead := &ListNode{}
	//l := dummyHead
	t1 := l1
	t2 := l2

	n1 := t1.listNodeToNum()
	n2 := t2.listNodeToNum()

	return makeListNode(n1 + n2)
}

func main() {
	l1 := makeListNode(342)
	l2 := makeListNode(465)
	fmt.Println(l1.listNodeToNum())
	fmt.Println(l2.listNodeToNum())
	l1.printListNode()
	l2.printListNode()
	l := addTwoNumbers(l1, l2)
	if l != nil {
		l.printListNode()
	}
}
