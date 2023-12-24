package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
)

// Clone 克隆
func (self PriorityQueue[T]) Clone() PriorityQueue[T] {
	return PriorityQueue[T](stack.Heap[node[T]](self).Clone())
}
