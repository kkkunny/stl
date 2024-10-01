//go:build goexperiment.rangefunc || go1.23

package linkedhashmap

import "iter"

type linkedhashmapIter[K, V any] interface {
	Iter2() iter.Seq2[K, V]
}
