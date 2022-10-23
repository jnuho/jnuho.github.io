package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key byte
	left *Node
	right *Node
}

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{data}
	} else {
		t.root.insert(data)
	}
}

func (node *Node) insert(data byte) {
	if node == nil {
		node = &Node{data}
	} else {
		if node.data >= data {
			if node.left == nil {
				node.left = &Node{data}
			} else {
				node.left.insert(data)
			}
		} else {
			if node.right == nil {
				node.right = &Node{data}
			} else {
				node.right.insert(data)
			}
		}
	}
}


func main() {
	var t Tree
	t.insert('A')

}
