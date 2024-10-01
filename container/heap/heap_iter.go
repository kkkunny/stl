//go:build goexperiment.rangefunc || go1.23

package stlheap

import "iter"

type heapIter[T any] interface {
	Iter() iter.Seq[T]
}
