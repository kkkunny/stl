//go:build goexperiment.rangefunc || go1.23

package bimap

import "iter"

type bimapIter[T, E any] interface {
	Iter2() iter.Seq2[T, E]
}
