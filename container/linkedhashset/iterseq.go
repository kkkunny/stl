//go:build goexperiment.rangefunc

package linkedhashset

import "iter"

func (self LinkedHashSet[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
