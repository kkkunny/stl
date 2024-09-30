package heap

import (
	stlbasic "github.com/kkkunny/stl/cmp"
	stlslices "github.com/kkkunny/stl/container/slices"
)

func (self *Heap[T]) init() {
	if self.data != nil {
		return
	}
	self.data = make([]T, 0)
}

// 获取左子节点下标
func (self *Heap[T]) getLeftSonIndex(i uint) uint {
	return 2*i + 1
}

// 获取右子节点下标
func (self *Heap[T]) getRightSonIndex(i uint) uint {
	return 2*i + 2
}

// 获取父节点下标
func (self *Heap[T]) getParentIndex(i uint) uint {
	return (i - 1) / 2
}

// 比较并交换
func (self *Heap[T]) compareAndExchange(pi, si uint) bool {
	pv, sv := self.data[pi], self.data[si]
	sort := stlbasic.Compare(pv, sv)
	if !((!self.reverse && sort > 0) || (self.reverse && sort < 0)) {
		return false
	}
	self.data[pi] = sv
	self.data[si] = pv
	return true
}

// 从下标开始修复
func (self *Heap[T]) fix(i uint) {
	for i > 0 {
		pi := self.getParentIndex(i)
		self.compareAndExchange(pi, i)
		i = pi
	}
}

// Push 入堆 O(log(N))
func (self *Heap[T]) Push(v T, vs ...T) {
	self.init()
	for _, v = range append([]T{v}, vs...) {
		self.data = append(self.data, v)
		self.fix(uint(len(self.data)) - 1)
	}
}

// Pop 出堆 O(log(N))
func (self *Heap[T]) Pop() T {
	self.init()
	v := stlslices.First(self.data)
	if len(self.data) == 1 {
		self.data = self.data[1:]
	} else {
		self.data[0] = stlslices.Last(self.data)
		self.data = self.data[:len(self.data)-1]
		self.fix(uint(len(self.data)) - 1)
	}
	return v
}

// Peek 头节点 O(1)
func (self Heap[T]) Peek() T {
	self.init()
	return stlslices.First(self.data)
}

// Reverse 反转 O(log(N))
func (self *Heap[T]) Reverse() {
	self.init()
	self.reverse = !self.reverse
	self.fix(self.Length() - 1)
}

// Clear 清空
func (self *Heap[T]) Clear() {
	self.data = nil
}

// Empty 是否为空
func (self Heap[T]) Empty() bool {
	return self.data == nil || stlslices.Empty(self.data)
}
