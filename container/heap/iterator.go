package heap

import (
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	stliter "github.com/kkkunny/stl/container/iter"
)

func (self Heap[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	return Heap[T]{
		reverse: false,
		data:    stlbasic.Ptr(dynarray.DynArray[T]{}.NewWithIterator(iter).(dynarray.DynArray[T])),
	}
}

// Iterator 迭代器
func (self Heap[T]) Iterator() stliter.Iterator[T] {
	self.init()
	data := self.data.Clone()
	data.Sort(self.reverse)
	return data.Iterator()
}
