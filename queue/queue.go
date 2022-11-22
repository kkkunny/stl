package queue

import (
	"github.com/kkkunny/stl/list"
)

// Queue 队列
type Queue[T any] struct {
	data *list.SingleLinkedList[T]
}

// NewQueue 新建队列
func NewQueue[T any](e ...T) *Queue[T] {
	return &Queue[T]{
		data: list.NewSingleLinkedList(e...),
	}
}

// 转成字符串 O(N)
func (self *Queue[T]) String() string {
	return self.data.String()
}

// Length 获取长度 O(1)
func (self *Queue[T]) Length() int {
	return self.data.Length()
}

// Empty 是否为空 O(1)
func (self *Queue[T]) Empty() bool {
	return self.data.Empty()
}

// Push 压入队列 O(1)
func (self *Queue[T]) Push(e ...T) {
	self.data.PushBack(e...)
}

// Pop 弹出队列 O(1)
func (self *Queue[T]) Pop() T {
	return self.data.PopFront()
}

// Peek 获取队首 O(1)
func (self *Queue[T]) Peek() T {
	return self.data.First()
}

// Clear 清空 O(1)
func (self *Queue[T]) Clear() {
	self.data.Clear()
}

// Clone 克隆 O(N)
func (self *Queue[T]) Clone() *Queue[T] {
	return &Queue[T]{data: self.data.Clone()}
}
