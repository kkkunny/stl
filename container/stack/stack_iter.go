//go:build goexperiment.rangefunc || go1.23

package stack

import "iter"

type stackIter[T any] interface {
	Iter() iter.Seq[T]
}

func (self *_Stack[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(self.data) - 1; i >= 0; i-- {
			if !yield(self.data[i]) {
				return
			}
		}
	}
}
