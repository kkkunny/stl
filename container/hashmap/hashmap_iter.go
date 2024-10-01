//go:build goexperiment.rangefunc || go1.23

package hashmap

import "iter"

type hashmapIter2[K, V any] interface {
	Iter2() iter.Seq2[K, V]
}
