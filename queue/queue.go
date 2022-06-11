package queue

import (
	"fmt"
	"github.com/kkkunny/stl/list"
)

// 队列
type Queue[T any] struct {
	data *list.SingleLinkedList[T]
}

// 新建队列
func NewQueue[T any](e ...T) *Queue[T] {
	return &Queue[T]{
		data: list.NewSingleLinkedList(e...),
	}
}

// 转成字符串 O(N)
func (self *Queue[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度 O(1)
func (self *Queue[T]) Length() int {
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

// 获取队首 O(1)
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
