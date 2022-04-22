package stack

import (
	"fmt"
)

// 栈
type Stack[T any] struct {
	data []T
}

// 新建栈
func NewStack[T any](e ...T) *Stack[T] {
	return &Stack[T]{
		data: e,
	}
}

// 转成字符串 O(N)
func (self *Stack[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度 O(1)
func (self *Stack[T]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *Stack[T]) Empty() bool {
	return len(self.data) == 0
}

// 压入栈 O(1)
func (self *Stack[T]) Push(e ...T) {
	self.data = append(self.data, e...)
}

// 弹出栈 O(1)
func (self *Stack[T]) Pop() T {
	v := self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	return v
}

// 获取栈顶 O(1)
func (self *Stack[T]) Peek() T {
	return self.data[len(self.data)-1]
}

// 清空 O(1)
func (self *Stack[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// 克隆 O(N)
func (self *Stack[T]) Clone() *Stack[T] {
	data := make([]T, len(self.data))
	copy(data, self.data)
	return &Stack[T]{data: data}
}
