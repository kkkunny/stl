package heap

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Clone 克隆
func (self Heap[T]) Clone() Heap[T] {
	self.init()
	return Heap[T]{
		reverse: self.reverse,
		data:    stlslices.Clone(self.data),
	}
}
