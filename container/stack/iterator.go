package stack

import (
	"slices"

	"github.com/kkkunny/stl/container/dynarray"
	stliter "github.com/kkkunny/stl/container/iter"
)

func (self Stack[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	return Stack[T](dynarray.DynArray[T](self).NewWithIterator(iter).(dynarray.DynArray[T]))
}

// Iterator 迭代器
func (self Stack[T]) Iterator() stliter.Iterator[T] {
	reverse := slices.Clone(dynarray.DynArray[T](self).ToSlice())
	slices.Reverse(reverse)
	return dynarray.NewDynArrayWith(reverse...).Iterator()
}
