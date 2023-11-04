package iterator

import stlbasic "github.com/kkkunny/stl/basic"

// 迭代器
type _Iter[T any] interface {
	stlbasic.Lengthable
	Next() bool
	HasNext() bool
	Value() T
	Reset()
}

// IteratorContainer 迭代器容器
type IteratorContainer[T any] interface {
	AppendWithIterator(iter Iterator[T])
	Append(ctr IteratorContainer[T])
	Iterator() Iterator[T]
}
