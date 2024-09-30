package queue

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Push 入队
func (self *Queue[T]) Push(v T, vs ...T) {
	*self = append(*self, v)
	*self = append(*self, vs...)
}

// Pop 出队
func (self *Queue[T]) Pop() T {
	v := stlslices.First(*self)
	*self = (*self)[1:]
	return v
}

// Peek 头元素
func (self Queue[T]) Peek() T {
	return stlslices.First(self)
}

// Clear 清空
func (self *Queue[T]) Clear() {
	*self = make([]T, 0)
}

// Empty 是否为空
func (self Queue[T]) Empty() bool {
	return stlslices.Empty(self)
}

// ToSlice 转成切片
func (self Queue[T]) ToSlice() []T {
	return self
}
