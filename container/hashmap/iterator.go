package hashmap

import (
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

func (_ HashMap[K, V]) NewWithIterator(iter iterator.Iterator[pair.Pair[K, V]]) any {
	self := NewHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self HashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	self.init()
	return self.KeyValues().Iterator()
}
