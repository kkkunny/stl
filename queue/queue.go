package queue

import (
	"fmt"
)

// 队列
type Queue[T any] struct {
	data []T
}

// 新建队列
func NewQueue[T any](e ...T) *Queue[T] {
	return &Queue[T]{
		data: e,
	}
}

// 转成字符串 O(N)
func (self *Queue[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度 O(1)
func (self *Queue[T]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *Queue[T]) Empty() bool {
	return len(self.data) == 0
}

// 压入队列 O(1)
func (self *Queue[T]) Push(e ...T) {
	self.data = append(self.data, e...)
}

// 弹出队列 O(1)
func (self *Queue[T]) Pop() T {
	v := self.data[0]
	self.data = self.data[1:]
	return v
}

// 获取队首 O(1)
func (self *Queue[T]) Peek() T {
	return self.data[0]
}

// 清空 O(1)
func (self *Queue[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// 克隆 O(N)
func (self *Queue[T]) Clone() *Queue[T] {
	data := make([]T, len(self.data))
	copy(data, self.data)
	return &Queue[T]{data: data}
}
