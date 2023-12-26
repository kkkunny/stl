package treemap

import (
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

func (_ TreeMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := NewTreeMap[K, V]()
	for iter.Next() {
		node := iter.Value()
		self.Set(node.First, node.Second)
	}
	return self
}

func (self TreeMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	self.init()
	return self.KeyValues().Iterator()
}
