//go:build goexperiment.rangefunc

package queue

import "iter"

func (self *_Queue[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
