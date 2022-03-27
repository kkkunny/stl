package stack

import (
	"stl/list"
	. "stl/types"
)

// 栈
type Stack[T any] struct {
	data *list.LinkedList[T]
}

// 新建栈
func NewStack[T any](e ...T) *Stack[T] {
	return &Stack[T]{
		data: list.NewLinkedList(e...),
	}
}

// 转成字符串
func (self *Stack[T]) String() string {
	return self.data.String()
}

// 获取长度
func (self *Stack[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空
func (self *Stack[T]) Empty() bool {
	return self.data.Empty()
}

// 压入栈
func (self *Stack[T]) Push(e ...T) {
	self.data.PushBack(e...)
}

// 弹出栈
func (self *Stack[T]) Pop() T {
	return self.data.PopBack()
}

// 清空
func (self *Stack[T]) Clear() {
	self.data.Clear()
}

// 克隆
func (self *Stack[T]) Clone() *Stack[T] {
	return &Stack[T]{
		data: self.data.Clone(),
	}
}

// 获取迭代器
func (self *Stack[T]) Iterator() *list.LinkedListIterator[T] {
	return self.data.Begin()
}
