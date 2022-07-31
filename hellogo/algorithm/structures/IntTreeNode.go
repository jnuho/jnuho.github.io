package structures

// binary tree data structure
type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

// node in a singly linked list

type ListNode struct {
	Val int
	Next *ListNode
}