package iterator

// Iterator 迭代器
type Iterator[Ctr, V any] struct {
	_iter[V]
}

func NewIterator[Ctr, V any](iter _iter[V]) Iterator[Ctr, V] {
	return Iterator[Ctr, V]{_iter: iter}
}

func (self Iterator[Ctr, V]) Collect() Ctr {
	var ctr Ctr
	ctr = any(ctr).(_iterContainer[Ctr, V]).NewWithIterator(self)
	return ctr
}
