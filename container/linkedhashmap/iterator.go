package linkedhashmap

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

func (_ LinkedHashMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := NewLinkedHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self LinkedHashMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	self.init()
	return stliter.NewSliceIterator(self.KeyValues()...)
}
