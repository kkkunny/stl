//go:build goexperiment.rangefunc || go1.23

package set

import "iter"

type setIter[T any] interface {
	Iter() iter.Seq[T]
}

func (self *_Set[T]) Iter() iter.Seq[T] {
	f := self.data.Iter2()
	return func(yield func(T) bool) {
		f(func(v T, _ struct{}) bool {
			return yield(v)
		})
	}
}
