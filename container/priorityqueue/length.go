package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
)

// Length 长度
func (self PriorityQueue[T]) Length() uint {
	return stack.Heap[node[T]](self).Length()
}
