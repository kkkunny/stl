//go:build goexperiment.rangefunc

package linkedlist

import "iter"

func (self LinkedList[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
