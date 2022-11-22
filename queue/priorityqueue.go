package queue

import (
	"github.com/kkkunny/stl/heap"
	"golang.org/x/exp/constraints"
)

// 优先级队列元素
type priorityQueueElem[P constraints.Ordered, V any] struct {
	Priority P
	Value    V
}

func (self priorityQueueElem[P, V]) Compare(dst priorityQueueElem[P, V]) int {
	if self.Priority > dst.Priority {
		return -1
	} else if self.Priority == dst.Priority {
		return 0
	} else {
		return 1
	}
}

// PriorityQueue 优先级队列
type PriorityQueue[P constraints.Ordered, V any] struct {
	data *heap.Container[priorityQueueElem[P, V]]
}

// NewPriorityQueue 新建优先级队列
func NewPriorityQueue[P constraints.Ordered, V any]() *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{data: heap.NewContainer[priorityQueueElem[P, V]]()}
}

// Length 获取长度 O(1)
func (self *PriorityQueue[P, V]) Length() int {
	return self.data.Length()
}

// Empty 是否为空 O(1)
func (self *PriorityQueue[P, V]) Empty() bool {
	return self.data.Empty()
}

// Push 压入队列 O(logN)-O(NlogN)
func (self *PriorityQueue[P, V]) Push(p P, v V) {
	self.data.Push(priorityQueueElem[P, V]{
		Priority: p,
		Value:    v,
	})
}

// Pop 弹出队列 O(NlogN)-O(N²logN)
func (self *PriorityQueue[P, V]) Pop() (P, V) {
	top := self.data.Pop()
	return top.Priority, top.Value
}

// Peek 获取队首 O(1)
func (self *PriorityQueue[P, V]) Peek() (P, V) {
	top := self.data.Peek()
	return top.Priority, top.Value
}

// Clear 清空 O(1)
func (self *PriorityQueue[P, V]) Clear() {
	if self.Empty() {
		return
	}
	self.data.Clear()
}

// Clone 克隆 O(N)
func (self *PriorityQueue[P, V]) Clone() *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{data: self.data.Clone()}
}
