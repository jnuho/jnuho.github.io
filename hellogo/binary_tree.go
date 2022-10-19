package main
import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	key byte
	left *Node
	right *Node
}

func (t *Tree) insert(data byte) {
	if t == nil {
	} else {
	}


}

func (n *Node) insert(data byte) {
	if n.key <= data {
		if n == nil {
			n = &Node{key:data, left:nil, right:nil}
		} else {
			n.right.insert(data)
		}
	} else {
		if n == nil {
			n = &Node{key:data, left:nil, right:nil}
		} else {
			n.left.insert(data)
		}
	}
}

func main() {
	var t Tree
}
