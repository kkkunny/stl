//go:build goexperiment.rangefunc || go1.23

package pqueue

import "iter"

func (self *_AnyPQueue[T]) Iter2() iter.Seq2[T, uint64] {
	f := self.data.Iter()
	return func(yield func(T, uint64) bool) {
		f(func(node anyPQueueNode[T]) bool {
			return yield(node.value, node.priority)
		})
	}
}
