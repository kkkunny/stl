//go:build !goexperiment.rangefunc && !go1.23

package stliter

// IterContainer 迭代器容器
type IterContainer[T any] interface {
	NewWithIterator(iter Iterator[T]) any
	Iterator() Iterator[T]
}

type Iter2Container[T any] interface {
	IterContainer[T]
}
