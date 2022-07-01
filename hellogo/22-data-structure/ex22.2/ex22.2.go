package main

import (
	"fmt"
	"container/list"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueu() *Queue {
	return &Queue{ list.New()}
}

func main() {
	queue := NewQueue()

	for i :=1; i<=4; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Print(v, " ")
		v = queue.Pop()
	}
	fmt.Println()
}