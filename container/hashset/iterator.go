package hashset

import (
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
)

func (_ HashSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := NewHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self HashSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(hashmap.HashMap[T, struct{}](self).Keys()...)
}
