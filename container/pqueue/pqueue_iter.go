//go:build goexperiment.rangefunc || go1.23

package pqueue

import "iter"

type pqueueIter[T any] interface {
	Iter2() iter.Seq2[T, uint64]
}
