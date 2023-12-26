//go:build goexperiment.rangefunc

package dynarray

import "iter"

func (self DynArray[T]) IterSeq() iter.Seq2[int, T] {
	self.init()
	return func(yield func(int, T) bool) {
		for i, e := range *self.data {
			if !yield(i, e) {
				return
			}
		}
	}
}
