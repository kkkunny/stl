//go:build goexperiment.rangefunc || go1.23

package treeset

import "iter"

type treesetIter[T any] interface {
	Iter() iter.Seq[T]
}
