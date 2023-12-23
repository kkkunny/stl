package stack

import (
	"github.com/kkkunny/stl/container/dynarray"
)

// Push 入队
func (self *Stack[T]) Push(v T, vs ...T) {
	(*dynarray.DynArray[T])(self).PushBack(v, vs...)
}

// Pop 出队
func (self *Stack[T]) Pop() T {
	return (*dynarray.DynArray[T])(self).PopBack()
}

// Peek 头元素
func (self Stack[T]) Peek() T {
	return dynarray.DynArray[T](self).Back()
}

// Clear 清空
func (self *Stack[T]) Clear() {
	(*dynarray.DynArray[T])(self).Clear()
}

// Empty 是否为空
func (self Stack[T]) Empty() bool {
	return dynarray.DynArray[T](self).Empty()
}

// ToSlice 转成切片
func (self Stack[T]) ToSlice() []T {
	return dynarray.DynArray[T](self).ToSlice()
}
