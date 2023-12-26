package queue

import (
	"github.com/kkkunny/stl/container/dynarray"
	stliter "github.com/kkkunny/stl/container/iter"
)

func (self Queue[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	return Queue[T](dynarray.DynArray[T](self).NewWithIterator(iter).(dynarray.DynArray[T]))
}

// Iterator 迭代器
func (self Queue[T]) Iterator() stliter.Iterator[T] {
	return dynarray.DynArray[T](self).Iterator()
}
