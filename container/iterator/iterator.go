package iterator

// Iterator 迭代器
type Iterator[Ctr, V any] struct {
	_Iter[V]
}

func NewIterator[Ctr, V any](iter _Iter[V]) Iterator[Ctr, V] {
	return Iterator[Ctr, V]{_Iter: iter}
}

func (self Iterator[Ctr, V]) Collect() Ctr {
	var ctr Ctr
	ctr = any(ctr).(_IterContainer[Ctr, V]).NewWithIterator(self)
	return ctr
}

func (self Iterator[Ctr, V]) Foreach(f func(v V)) {
	for self.Next() {
		f(self.Value())
	}
}

func (self Iterator[Ctr, V]) Map(f func(v V) V) Ctr {
	vs := make([]V, self.Length())
	var i int
	self.Foreach(func(v V) {
		vs[i] = f(v)
		i++
	})
	var ctr Ctr
	ctr = any(ctr).(_IterContainer[Ctr, V]).NewWithIterator(NewIterator[Ctr, V](_NewSliceIterator[V](vs...)))
	return ctr
}
