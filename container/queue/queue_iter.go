//go:build goexperiment.rangefunc || go1.23

package queue

import (
	"iter"

	stliter "github.com/kkkunny/stl/container/iter"
)

type queueIter[T any] interface {
	Iter() iter.Seq[T]
}

func (*_Queue[T]) FromIter(seq iter.Seq[T]) Queue[T] {
	return _NewQueueWith(stliter.Collect[T, []T](seq)...)
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
