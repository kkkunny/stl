package stack

import stlbasic "github.com/kkkunny/stl/basic"

// Clone 克隆
func (self Heap[T]) Clone() Heap[T] {
	self.init()
	return Heap[T]{
		reverse: self.reverse,
		data:    stlbasic.Ptr(self.data.Clone()),
	}
}
