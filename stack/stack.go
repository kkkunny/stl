package stack

import (
	"github.com/kkkunny/stl/list"
)

// 栈
type Stack[T any] struct {
	data *list.SingleLinkedList[T]
}

// 新建栈
func NewStack[T any](e ...T) *Stack[T] {
	return &Stack[T]{
		data: list.NewSingleLinkedList[T](e...),
	}
}

// 转成字符串 O(N)
func (self *Stack[T]) String() string {
	return self.data.String()
}

// 获取长度 O(1)
func (self *Stack[T]) Length() int {
	return self.data.Length()
}

// 是否为空 O(1)
func (self *Stack[T]) Empty() bool {
	return self.data.Empty()
}

// 压入栈 O(1)
func (self *Stack[T]) Push(e ...T) {
	self.data.PushBack(e...)
}

// 弹出栈 O(1)
func (self *Stack[T]) Pop() T {
	return self.data.PopBack()
}

// 获取栈顶 O(1)
func (self *Stack[T]) Peek() T {
	return self.data.Last()
}

// 清空 O(1)
func (self *Stack[T]) Clear() {
	self.data.Clear()
}

// 克隆 O(N)
func (self *Stack[T]) Clone() *Stack[T] {
	return &Stack[T]{data: self.data.Clone()}
}
