package iterator

// IteratorContainer 迭代器容器
type IteratorContainer[T any] interface {
	AppendWithIterator(iter Iterator[T])
	Append(ctr IteratorContainer[T])
	Iterator() Iterator[T]
}
