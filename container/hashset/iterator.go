package hashset

import (
	"github.com/kkkunny/stl/container/iterator"
)

func (_ HashSet[T]) NewWithIterator(iter iterator.Iterator[T]) any {
	self := NewHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Push(iter.Value())
	}
	return self
}

func (self HashSet[T]) Iterator() iterator.Iterator[T] {
	return self.data.Keys().Iterator()
}
