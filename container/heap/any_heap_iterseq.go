//go:build goexperiment.rangefunc

package stlheap

import "iter"

func (self *_AnyHeap[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
