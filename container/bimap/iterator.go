package bimap

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

func (_ BiMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := NewBiMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self BiMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}
