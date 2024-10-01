//go:build goexperiment.rangefunc || go1.23

package treemap

import "iter"

type treemapIter[K, V any] interface {
	Iter2() iter.Seq2[K, V]
}
