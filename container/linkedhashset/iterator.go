package linkedhashset

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

func (_ LinkedHashSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := NewLinkedHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self LinkedHashSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(linkedhashmap.LinkedHashMap[T, struct{}](self).Keys()...)
}
