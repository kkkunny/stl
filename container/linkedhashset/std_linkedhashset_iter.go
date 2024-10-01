//go:build goexperiment.rangefunc || go1.23

package linkedhashset

import "iter"

func (self *_StdLinkedHashSet[T]) Iter() iter.Seq[T] {
	f := self.data.Iter2()
	return func(yield func(T) bool) {
		f(func(v T, _ struct{}) bool {
			return yield(v)
		})
	}
}
