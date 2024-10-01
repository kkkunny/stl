//go:build !goexperiment.rangefunc && !go1.23

package set

type setIter[T any] interface{}
