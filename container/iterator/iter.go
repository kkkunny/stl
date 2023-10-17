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

// 迭代器容器
type _IterContainer[Ctr, V any] interface {
    NewWithIterator(iter Iterator[V]) Ctr
    Iterator() Iterator[V]
}
