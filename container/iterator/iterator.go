package iterator

// Iterator 迭代器
type Iterator[V any] struct {
	_Iter[V]
}

func NewIterator[Ctr, V any](iter _Iter[V]) Iterator[V] {
	return Iterator[V]{_Iter: iter}
}

func (self Iterator[V]) Foreach(f func(v V)) {
	for self.Next() {
		f(self.Value())
	}
}
