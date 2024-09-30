package stack

import (
	"golang.org/x/exp/slices"

	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
)

func (self Stack[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return Stack[T](data)
}

// Iterator 迭代器
func (self Stack[T]) Iterator() stliter.Iterator[T] {
	reverse := stlslices.Clone(self)
	slices.Reverse(reverse)
	return stliter.NewSliceIterator(reverse...)
}
