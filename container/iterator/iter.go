package iterator

import stlbasic "github.com/kkkunny/stl/basic"

// 迭代器
type _Iter[T any] interface {
	stlbasic.Length
	Next() bool
	Value() T
	Reset()
}

// 迭代器容器
type _IterContainer[Ctr, V any] interface {
	NewWithIterator(iter Iterator[Ctr, V]) Ctr
	Iterator() Iterator[Ctr, V]
}
