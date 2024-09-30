package heap

import (
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
)

func (self Heap[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return Heap[T]{
		reverse: false,
		data:    data,
	}
}

// Iterator 迭代器
func (self Heap[T]) Iterator() stliter.Iterator[T] {
	self.init()
	data := stlslices.Sort(stlslices.Clone(self.data), self.reverse)
	return stliter.NewSliceIterator(data...)
}
