package stack

import (
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
)

func (self Heap[T]) NewWithIterator(iter iterator.Iterator[T]) any {
	return Heap[T]{
		reverse: false,
		data:    stlbasic.Ptr(dynarray.DynArray[T]{}.NewWithIterator(iter).(dynarray.DynArray[T])),
	}
}

// Iterator 迭代器
func (self Heap[T]) Iterator() iterator.Iterator[T] {
	self.init()
	return self.data.Iterator()
}
