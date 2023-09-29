package hashset

import (
	hashmap "github.com/kkkunny/stl/container/hash_map"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

type _iterator[T comparable] struct {
	iter iterator.Iterator[hashmap.HashMap[T, struct{}], pair.Pair[T, struct{}]]
}

func _NewIterator[T comparable](src *HashSet[T]) *_iterator[T] {
	return &_iterator[T]{iter: src.data.Iterator()}
}

func (self *_iterator[T]) Length() uint {
	return self.Length()
}

func (self *_iterator[T]) Next() bool {
	return self.Next()
}

func (self _iterator[T]) HasNext() bool {
	return self.HasNext()
}

func (self _iterator[T]) Value() T {
	return self.iter.Value().First
}

func (self *_iterator[T]) Reset() {
	self.iter.Reset()
}
