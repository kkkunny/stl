//go:build goexperiment.rangefunc

package pqueue

import "iter"

func (self *_AnyPQueue[T]) IterSeq() iter.Seq2[uint64, T] {
	return func(yield func(uint64, T) bool) {
		for i := self.Iterator(); i.Next(); {
			elem := i.Value()
			if !yield(elem.First, elem.Second) {
				return
			}
		}
	}
}
