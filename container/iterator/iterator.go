package iterator

// Iterator 迭代器
type Iterator[V any] struct {
	_Iter[V]
}

func NewIterator[V any, Iter _Iter[V]](iter Iter) Iterator[V] {
	return Iterator[V]{_Iter: iter}
}

func (self Iterator[V]) Foreach(f func(v V) bool) {
	for self.Next() {
		if !f(self.Value()) {
			break
		}
	}
}
