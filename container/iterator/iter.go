package iterator

// 迭代器
type _iter[T any] interface {
	Next() bool
	Value() T
}

// 迭代器容器
type _iterContainer[Ctr, V any] interface {
	NewWithIterator(iter Iterator[Ctr, V]) Ctr
	Iterator() Iterator[Ctr, V]
}
