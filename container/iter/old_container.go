package stliter

// IteratorContainer 迭代器容器
type IteratorContainer[T any] interface {
	NewWithIterator(iter Iterator[T]) any
	Iterator() Iterator[T]
}
