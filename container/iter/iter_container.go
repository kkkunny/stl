//go:build goexperiment.rangefunc || go1.23

package stliter

import "iter"

type IterContainer[T, Container any] interface {
	FromIter(seq iter.Seq[T]) Container
	Iter() iter.Seq[T]
}
