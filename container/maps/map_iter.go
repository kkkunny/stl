//go:build goexperiment.rangefunc || go1.23

package stlmaps

import "iter"

type mapIter2[K, V any] interface {
	Iter2() iter.Seq2[K, V]
}
