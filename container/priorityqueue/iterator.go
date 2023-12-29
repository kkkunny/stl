package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

func (self PriorityQueue[T]) NewWithIterator(iter stliter.Iterator[pair.Pair[uint64, T]]) any {
	pq := NewPriorityQueue[T]()
	stliter.IteratorForeach(iter, func(v pair.Pair[uint64, T]) bool {
		pq.Push(v.First, v.Second)
		return true
	})
	return pq
}

// Iterator 迭代器
func (self PriorityQueue[T]) Iterator() stliter.Iterator[pair.Pair[uint64, T]] {
	return stliter.IteratorMap(stack.Heap[node[T]](self).Iterator(), func(v node[T]) pair.Pair[uint64, T] {
		return pair.NewPair(v.priority, v.value)
	})
}
