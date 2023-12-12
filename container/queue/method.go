package queue

import (
	"github.com/kkkunny/stl/container/dynarray"
)

// Push 入队
func (self *Queue[T]) Push(v T, vs ...T) {
	(*dynarray.DynArray[T])(self).PushBack(v, vs...)
}

// Pop 出队
func (self *Queue[T]) Pop() T {
	return (*dynarray.DynArray[T])(self).PopFront()
}

// Peek 头元素
func (self Queue[T]) Peek() T {
	return dynarray.DynArray[T](self).Front()
}

// Clear 清空
func (self *Queue[T]) Clear() {
	(*dynarray.DynArray[T])(self).Clear()
}

// Empty 是否为空
func (self Queue[T]) Empty() bool {
	return dynarray.DynArray[T](self).Empty()
}

// Append 拼接
func (self *Queue[T]) Append(dst Queue[T]) {
	(*dynarray.DynArray[T])(self).Append(dynarray.DynArray[T](dst))
}

// ToSlice 转成切片
func (self Queue[T]) ToSlice() []T {
	return dynarray.DynArray[T](self).ToSlice()
}
