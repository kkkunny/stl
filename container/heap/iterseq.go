//go:build goexperiment.rangefunc

package heap

import "iter"

func (self Heap[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
