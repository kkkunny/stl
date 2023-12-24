package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
)

// Equal 比较相等
func (self PriorityQueue[T]) Equal(dst PriorityQueue[T]) bool {
	return stack.Heap[node[T]](self).Equal(stack.Heap[node[T]](dst))
}
