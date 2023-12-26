package treeset

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/treemap"
)

func (_ TreeSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := NewTreeSet[T]()
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self TreeSet[T]) Iterator() stliter.Iterator[T] {
	return treemap.TreeMap[T, struct{}](self).Keys().Iterator()
}
