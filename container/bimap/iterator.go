package bimap

import (
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

func (_ BiMap[K, V]) NewWithIterator(iter iterator.Iterator[pair.Pair[K, V]]) any {
	self := NewBiMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self BiMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return self.KeyValues().Iterator()
}
