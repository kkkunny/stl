package linkedhashset

import (
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

func (_ LinkedHashSet[T]) NewWithIterator(iter iterator.Iterator[T]) any {
	self := NewLinkedHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self LinkedHashSet[T]) Iterator() iterator.Iterator[T] {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).Keys().Iterator()
}
