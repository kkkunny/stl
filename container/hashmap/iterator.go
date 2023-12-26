package hashmap

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

func (_ HashMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := NewHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self HashMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	self.init()
	return self.KeyValues().Iterator()
}
