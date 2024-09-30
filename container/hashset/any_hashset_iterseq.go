//go:build goexperiment.rangefunc

package hashset

import "iter"

func (self *_AnyHashSet[T]) IterSeq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := self.Iterator(); i.Next(); {
			if !yield(i.Value()) {
				return
			}
		}
	}
}
