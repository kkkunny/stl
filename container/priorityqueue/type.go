package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
)

type node[T any] struct {
	priority uint64
	value    T
}

func (self node[T]) Equal(dst node[T]) bool {
	return self.priority == dst.priority
}

func (self node[T]) Compare(dst node[T]) int {
	if self.priority < dst.priority {
		return 1
	} else if self.priority == dst.priority {
		return 0
	} else {
		return -1
	}
}

// PriorityQueue 优先级队列
type PriorityQueue[T any] stack.Heap[node[T]]

// NewPriorityQueue 新建优先级队列
func NewPriorityQueue[T any]() PriorityQueue[T] {
	return PriorityQueue[T](stack.NewMinHeap[node[T]]())
}

func (self PriorityQueue[T]) Default() PriorityQueue[T] {
	return NewPriorityQueue[T]()
}
