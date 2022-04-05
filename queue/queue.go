package queue

import (
	"github.com/kkkunny/stl/list"
	. "github.com/kkkunny/stl/types"
)

// 队列
type Queue[T any] struct {
	data *list.LinkedList[T]
}

// 新建队列
func NewQueue[T any](e ...T) *Queue[T] {
	return &Queue[T]{
		data: list.NewLinkedList(e...),
	}
}

// 转成字符串 O(N)
func (self *Queue[T]) String() string {
	return self.data.String()
}

// 获取长度 O(1)
func (self *Queue[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空 O(1)
func (self *Queue[T]) Empty() bool {
	return self.data.Empty()
}

// 压入队列 O(1)
func (self *Queue[T]) Push(e ...T) {
	self.data.PushBack(e...)
}

// 弹出队列 O(1)
func (self *Queue[T]) Pop() T {
	return self.data.PopFront()
}

// 提前获取队首 O(1)
func (self *Queue[T]) Peek() T {
	return self.data.First()
}

// 清空 O(1)
func (self *Queue[T]) Clear() {
	self.data.Clear()
}

// 克隆 O(N)
func (self *Queue[T]) Clone() *Queue[T] {
	return &Queue[T]{
		data: self.data.Clone(),
	}
}

// 获取迭代器
func (self *Queue[T]) Iterator() *list.LinkedListIterator[T] {
	return self.data.Begin()
}
