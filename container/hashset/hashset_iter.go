//go:build goexperiment.rangefunc || go1.23

package hashset

import "iter"

type hashsetIter[T any] interface {
	Iter() iter.Seq[T]
}
