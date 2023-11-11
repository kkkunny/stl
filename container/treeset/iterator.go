package treeset

import (
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/treemap"
)

func (_ TreeSet[T]) NewWithIterator(iter iterator.Iterator[T]) any {
	self := NewTreeSet[T]()
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self TreeSet[T]) Iterator() iterator.Iterator[T] {
	return treemap.TreeMap[T, struct{}](self).Keys().Iterator()
}
