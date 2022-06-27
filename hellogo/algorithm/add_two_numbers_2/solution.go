package main

import (
	"fmt"
	"github.com/austingebauer/go-leetcode/structures"
	"math/big"
)


/**
cd goproject/algorithm/add_two_numbers_2
go mod init algorithm/add_two_numbers_2
go mod tidy
*/

func addTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {

	return nil
}

// Input: pointer to ListNode
// Output:  pointer to big.Int
func deptFirstNum(l *structures.ListNode) *big.Int {
	if l == nil {
		return big.NewInt(0)
	}

	v := deptFirstNum(l.Next)

	return v.Mul(v, big.NewInt(10)).Add(v, big.NewInt(int64(l.Val)))
}



func printListNode(l *structures.ListNode) {
	next := l.Next
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}

func NumberToLinkedList(n int) *structures.ListNode {
	return nil
}

func NodeToNumber(l *structures.ListNode) int {
	next := l.Next
	multi := 1
	res := 0
	for next != nil {

		res += (multi * next.Val)
		next = next.Next
		multi *= 10
	}
	return res
}

func main() {
	// 256 +
	// l1 = [2,4,3]
	// l2 = [5,6,4]
	// Output: [7,0,8] Explanation: 342 + 465 = 807.
	dummyHead1 := &structures.ListNode{0, nil}
	dummyHead1.Next = &structures.ListNode{2, nil}
	dummyHead1.Next.Next = &structures.ListNode{4, nil}
	dummyHead1.Next.Next.Next = &structures.ListNode{3, nil}

	n := NodeToNumber(dummyHead1)
	fmt.Println(n)


	printListNode(dummyHead1)
	fmt.Println(deptFirstNum(dummyHead1.Next))
}