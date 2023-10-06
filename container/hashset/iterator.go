package hashset

import (
    "github.com/kkkunny/stl/container/iterator"
    "github.com/kkkunny/stl/container/pair"
)

type _iterator[T any] struct {
    iter iterator.Iterator[pair.Pair[T, struct{}]]
}

func _NewIterator[T any](src *HashSet[T]) *_iterator[T] {
    return &_iterator[T]{iter: src.data.Iterator()}
}

func (self _iterator[T]) Length() uint {
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
