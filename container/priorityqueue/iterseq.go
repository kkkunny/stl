//go:build goexperiment.rangefunc

package queue

import "iter"

func (self PriorityQueue[T]) IterSeq() iter.Seq2[uint64, T] {
	return func(yield func(uint64, T) bool) {
		for i := self.Iterator(); i.Next(); {
			elem := i.Value()
			if !yield(elem.First, elem.Second) {
				return
			}
		}
	}
}
