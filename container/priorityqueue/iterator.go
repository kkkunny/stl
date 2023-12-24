package queue

import (
	stack "github.com/kkkunny/stl/container/heap"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

func (self PriorityQueue[T]) NewWithIterator(iter iterator.Iterator[pair.Pair[uint64, T]]) any {
	pq := NewPriorityQueue[T]()
	iterator.IteratorForeach(iter, func(v pair.Pair[uint64, T]) bool {
		pq.Push(v.First, v.Second)
		return true
	})
	return pq
}

// Iterator 迭代器
func (self PriorityQueue[T]) Iterator() iterator.Iterator[pair.Pair[uint64, T]] {
	return iterator.IteratorMap(stack.Heap[node[T]](self).Iterator(), func(v node[T]) pair.Pair[uint64, T] {
		return pair.NewPair(v.priority, v.value)
	})
}
