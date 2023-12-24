package stack

import (
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
)

func (self *Heap[T]) init() {
	if self.data != nil {
		return
	}
	self.data = stlbasic.Ptr(dynarray.NewDynArray[T]())
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
	pv, sv := self.data.Get(pi), self.data.Get(si)
	sort := stlbasic.Order(pv, sv)
	if !((!self.reverse && sort > 0) || (self.reverse && sort < 0)) {
		return false
	}
	self.data.Set(pi, sv)
	self.data.Set(si, pv)
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
		self.data.PushBack(v)
		self.fix(self.data.Length() - 1)
	}
}

// Pop 出堆 O(log(N))
func (self *Heap[T]) Pop() T {
	self.init()
	v := self.data.Front()
	if self.data.Length() == 1 {
		self.data.PopFront()
	} else {
		self.data.Set(0, self.data.Back())
		self.data.PopBack()
		self.fix(self.data.Length() - 1)
	}
	return v
}

// Peek 头节点 O(1)
func (self Heap[T]) Peek() T {
	self.init()
	return self.data.Front()
}

// Reverse 反转 O(log(N))
func (self *Heap[T]) Reverse() {
	self.init()
	self.reverse = !self.reverse
	self.fix(self.Length() - 1)
}
