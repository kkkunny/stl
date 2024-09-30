package stack

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Push 入队
func (self *Stack[T]) Push(v T, vs ...T) {
	*self = append(*self, v)
	*self = append(*self, vs...)
}

// Pop 出队
func (self *Stack[T]) Pop() T {
	v := stlslices.Last(*self)
	*self = (*self)[:len(*self)-1]
	return v
}

// Peek 头元素
func (self Stack[T]) Peek() T {
	return stlslices.Last(self)
}

// Clear 清空
func (self *Stack[T]) Clear() {
	*self = make([]T, 0)
}

// Empty 是否为空
func (self Stack[T]) Empty() bool {
	return stlslices.Empty(self)
}

// ToSlice 转成切片
func (self Stack[T]) ToSlice() []T {
	return self
}
