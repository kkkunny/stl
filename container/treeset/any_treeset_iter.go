//go:build goexperiment.rangefunc || go1.23

package treeset

import "iter"

func (self *_AnyTreeSet[T]) Iter() iter.Seq[T] {
	f := self.data.Iter2()
	return func(yield func(T) bool) {
		f(func(v T, _ struct{}) bool {
			return yield(v)
		})
	}
}
