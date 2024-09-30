package queue

import (
	stliter "github.com/kkkunny/stl/container/iter"
)

func (self Queue[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return Queue[T](data)
}

// Iterator 迭代器
func (self Queue[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self...)
}
