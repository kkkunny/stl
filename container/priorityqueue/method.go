package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
)

// Push 入队
func (self *PriorityQueue[T]) Push(prior uint64, value T) {
	(*stack.Heap[node[T]])(self).Push(node[T]{
		priority: prior,
		value:    value,
	})
}

// Pop 出队
func (self *PriorityQueue[T]) Pop() (uint64, T) {
	node := (*stack.Heap[node[T]])(self).Pop()
	return node.priority, node.value
}

// Peek 头元素
func (self PriorityQueue[T]) Peek() (uint64, T) {
	node := stack.Heap[node[T]](self).Peek()
	return node.priority, node.value
}

// Clear 清空
func (self *PriorityQueue[T]) Clear() {
	(*stack.Heap[node[T]])(self).Clear()
}

// Empty 是否为空
func (self PriorityQueue[T]) Empty() bool {
	return stack.Heap[node[T]](self).Empty()
}
