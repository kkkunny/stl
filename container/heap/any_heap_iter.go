//go:build goexperiment.rangefunc || go1.23

package stlheap

import "iter"

func (self *_AnyHeap[T]) Iter() iter.Seq[T] {
	h := toAnyHeapData(self.data)
	return func(yield func(T) bool) {
		for _, v := range h.data {
			if !yield(v) {
				return
			}
		}
	}
}
