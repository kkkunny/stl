//go:build goexperiment.rangefunc || go1.23

package queue

import "iter"

type queueIter[T any] interface {
	Iter() iter.Seq[T]
}

func (self *_Queue[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range self.data {
			if !yield(v) {
				return
			}
		}
	}
}
