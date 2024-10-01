//go:build goexperiment.rangefunc || go1.23

package linkedhashset

import "iter"

type linkedhashsetIter[T any] interface {
	Iter() iter.Seq[T]
}
