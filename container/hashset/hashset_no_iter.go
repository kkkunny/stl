//go:build !goexperiment.rangefunc && !go1.23

package hashset

type hashsetIter[T any] interface{}
